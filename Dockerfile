# Простой Dockerfile без лишних шагов
FROM golang:1.21-alpine

WORKDIR /app

# Копируем все файлы проекта
COPY . .

# Собираем приложение
RUN go build -o tracker-app .

# Открываем порт
EXPOSE 8080

# Запускаем приложение
CMD ["./tracker-app"]
