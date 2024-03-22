# Стадия сборки для auth-service
FROM golang:1.22.1-alpine AS auth-builder
WORKDIR /build
COPY go.mod go.sum ./
# Установка git и других зависимостей необходимых для go mod tidy и go build
RUN apk add --no-cache git
RUN go mod download
COPY services/auth/ .
RUN go mod tidy
RUN go build -o auth-service .

# Стадия сборки для jwt-service
FROM golang:1.22.1-alpine AS jwt-builder
WORKDIR /build
COPY go.mod go.sum ./
# Установка git и других зависимостей необходимых для go mod tidy и go build
RUN apk add --no-cache git
RUN go mod download
COPY services/JWT/ .
RUN go mod tidy
RUN go build -o jwt-service .
