# Используем официальный базовый образ Go
FROM golang:1.16-alpine

# Установка рабочей директории внутри контейнера
WORKDIR /app

# Копируем go.mod и go.sum в рабочую директорию
COPY go.mod go.sum ./

# Загрузка зависимостей
RUN go mod download

# Копируем исходный код в контейнер
COPY . .

RUN go build -o /auth-service

CMD ["/auth-service"]