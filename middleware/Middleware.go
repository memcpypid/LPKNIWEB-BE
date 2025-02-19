package middleware

import (
	"LPKNI/lpkniService/config"
	"LPKNI/lpkniService/models"
	"fmt"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func VerifyJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ambil token JWT dari cookie
		tokenString, err := c.Cookie(os.Getenv("COOKIE_NAME"))
		if err != nil || tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Parsing token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Validasi algoritma
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			// Kembalikan kunci rahasia untuk validasi
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Ambil ID user dan session_id dari token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || claims["sub"] == nil || claims["session_id"] == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		userID := uint(claims["sub"].(float64))
		sessionID := claims["session_id"].(string)

		// Cari session berdasarkan userID dan sessionID
		var session models.SessionLogin
		if err := config.DB.Where("user_id = ? AND session_id = ?", userID, sessionID).First(&session).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid session"})
			c.Abort()
			return
		}

		// Cari data user dari userID
		var user models.AkunAnggota
		var DataAnggota models.DataAnggota
		if err := config.DB.First(&user, userID).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			c.Abort()
			return
		}
		if err := config.DB.Preload("ImageUsers").Preload("Wilayah").Preload("Daerah").Preload("JabatanStruktural").Where(models.DataAnggota{UserID: uint(userID)}).First(&DataAnggota).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "data Anggota not found"})
			c.Abort()
			return
		}
		// Set user dalam konteks request
		c.Set("user", user)
		c.Set("data_anggota", DataAnggota)
		c.Next()
	}
}
