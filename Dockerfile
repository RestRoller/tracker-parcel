FROM golang:1.21-alpine AS builder

WORKDIR /app

# Копируем все файлы проекта
COPY . .

# Собираем приложение (без отдельного go mod download)
RUN go build -o tracker-app .

# Финальный образ
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

# Копируем бинарник из builder
COPY --from=builder /app/tracker-app .

# Копируем базу данных (если есть)
COPY --from=builder /app/tracker.db ./

EXPOSE 8080

CMD ["./tracker-app"]
