# Gunakan image resmi Golang
FROM golang:1.20-alpine

# Set direktori kerja di dalam container
WORKDIR /app

# Salin go.mod dan go.sum ke dalam container
COPY go.mod go.sum ./

# Install dependencies
RUN go mod download

# Salin kode aplikasi ke dalam container
COPY . .

# Kompilasi aplikasi
RUN go build -o app .

# Jalankan aplikasi saat container dijalankan
CMD ["./app"]
