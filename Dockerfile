# Этап 1: Сборка приложения
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Копируем файлы модуля
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем приложение
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o tracker-app .

# Этап 2: Финальный образ
FROM alpine:latest

# Устанавливаем сертификаты для HTTPS
RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

# Копируем бинарный файл из этапа сборки
COPY --from=builder /app/tracker-app .

# Копируем базу данных (если нужна для инициализации)
COPY --from=builder /app/tracker.db ./

# Создаем volume для персистентных данных
VOLUME ["/app/data"]

# Открываем порт
EXPOSE 8080

# Запускаем приложение
CMD ["./tracker-app"]
