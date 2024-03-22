# Стадия сборки для auth-service
FROM golang:1.22.1-alpine AS auth-builder
WORKDIR /build
COPY go.mod go.sum ./
COPY services/auth/gen ./services/auth/gen
RUN go mod download
COPY services/auth/ .
RUN go mod tidy
RUN go build -o auth-service .

# Стадия сборки для jwt-service
FROM golang:1.22.1-alpine AS jwt-builder
WORKDIR /build
COPY go.mod go.sum ./
COPY services/JWT/gen ./services/JWT/gen
RUN go mod download
COPY services/JWT/ .
RUN go mod tidy
RUN go build -o jwt-service .

FROM golang:1.22.1-alpine AS rest-builder
WORKDIR /build
COPY go.mod go.sum ./
COPY frontend ./frontend
COPY restapi ./restapi
RUN go mod download
RUN go mod tidy
RUN go build -o rest-service .

FROM alpine
COPY --from=auth-builder /build/auth-service /auth-service
COPY --from=jwt-builder /build/jwt-service /jwt-service
COPY start-services.sh /start-services.sh
RUN chmod +x /start-services.sh
CMD ["/start-services.sh"]