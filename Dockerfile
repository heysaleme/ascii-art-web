# --- Stage 1: Build ---
FROM golang:1.24 AS builder

WORKDIR /app

# Модули
COPY go.mod ./
RUN go mod download || true

# Копируем весь проект
COPY . .

# Собираем приложение
RUN CGO_ENABLED=0 GOOS=linux go build -o server main.go web.go

# --- Stage 2: Run ---
FROM alpine:3.18

WORKDIR /app

# Копируем бинарник
COPY --from=builder /app/server .

# Копируем статику
COPY static ./static
COPY style ./style
COPY banners ./banners
COPY ascii ./ascii

EXPOSE 8080

CMD ["./server"]