# Этап сборки
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Копируем файлы модуля
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем приложение
RUN CGO_ENABLED=0 GOOS=linux go build -o tracker-app .

# Финальный образ
FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

# Копируем бинарник
COPY --from=builder /app/tracker-app .

# Копируем базу данных (если нужна)
COPY --from=builder /app/tracker.db ./

EXPOSE 8080

CMD ["./tracker-app"]
