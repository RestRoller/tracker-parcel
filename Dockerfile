FROM golang:1.24.3-alpine 

WORKDIR /app

# Копируем все файлы (ваши оригиналы)
COPY . .

# Собираем приложение
RUN go build -o app .

EXPOSE 8080

CMD ["./app"]
