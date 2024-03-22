# Стадия сборки для auth-service
FROM golang:1.12.1-alpine AS auth-builder
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY services/auth/ .
RUN go build -o auth-service .

# Стадия сборки для jwt-service
FROM golang:1.22.1-alpine AS jwt-builder
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY services/JWT/ .
RUN go build -o jwt-service .