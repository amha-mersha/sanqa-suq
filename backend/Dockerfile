FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o ./out/backend cmd/sanqa-suq-srv

FROM alpine:latest

WORKDIR /app
COPY --from=builder ./out/backend/ .
COPY --from=builder ./internal/database/migrations/ ./migrations

EXPOSE 8080

CMD ["./backend"]
