# Gunakan image Go resmi sebagai base image
FROM golang:1.21-alpine

# Set working directory di dalam container
WORKDIR /app

# Salin go.mod dan go.sum ke dalam container untuk mengelola dependencies
COPY go.mod go.sum ./

# Install dependencies Go
RUN go mod tidy

# Salin semua file aplikasi ke dalam container
COPY . .

# Build aplikasi Go
RUN go build -o app main.go

# Ekspos port yang digunakan oleh aplikasi (misalnya port 8080)
EXPOSE 8080

# Jalankan aplikasi
CMD ["./app", "--mode=redis"]