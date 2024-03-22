FROM golang:1.22.1-alpine AS auth-builder
WORKDIR /build
COPY go.mod go.sum ./
COPY services/auth/gen ./services/auth/gen
RUN go mod download
COPY services/auth/ .
RUN go mod tidy
RUN go build -o auth-service .

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
COPY frontend/templates ./frontend/templates
COPY frontend/styles ./frontend/styles
COPY frontend/scripts ./frontend/scripts
COPY restapi/router ./restapi/router
COPY services/auth/gen ./services/auth/gen
COPY restapi/cmd .
RUN go mod download
RUN go mod tidy
RUN go build -o rest-service .

COPY restapi/cmd .
RUN go mod download
RUN go mod tidy

FROM alpine
COPY --from=auth-builder /build/auth-service /auth-service
COPY --from=jwt-builder /build/jwt-service /jwt-service
COPY --from=rest-builder /build/frontend /frontend
COPY --from=rest-builder /build/services/auth/gen ./services/auth/gen
COPY --from=rest-builder /build/rest-service /rest-service
COPY start-services.sh /start-services.sh
RUN chmod +x /start-services.sh
CMD ["/start-services.sh"]