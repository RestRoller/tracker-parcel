FROM golang:1.21

WORKDIR /app

# Копируем все файлы проекта
COPY . .

# Устанавливаем зависимости (если есть go.mod)
RUN go mod download || echo "No dependencies to download"

# Собираем приложение с явным указанием имени выходного файла
RUN go build -o tracker-app -v .

# Указываем порт
EXPOSE 8080

# Запускаем приложение
CMD ["./tracker-app"]
