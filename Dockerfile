# Multi-stage build для Go приложения

# Stage 1: Build
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Устанавливаем зависимости для сборки
RUN apk add --no-cache git gcc musl-dev

# Копируем go mod и sum
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем приложение
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Stage 2: Run
FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/

# Копируем бинарник из builder
COPY --from=builder /app/main .

# Копируем базу данных (если нужно)
COPY --from=builder /app/tracker.db ./

# Создаем volume для данных
VOLUME ["/root/data"]

# Открываем порт
EXPOSE 8080

# Запускаем приложение
CMD ["./main"]