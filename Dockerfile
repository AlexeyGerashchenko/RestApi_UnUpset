# Используем официальный образ Go
FROM golang:1.23-alpine

# Устанавливаем необходимые инструменты
RUN apk add --no-cache gcc musl-dev

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы зависимостей
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем приложение
RUN go build -o main ./cmd

# Запускаем приложение
CMD ["./main"]
