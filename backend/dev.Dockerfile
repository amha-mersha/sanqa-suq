# Build stage
FROM golang:1.24-alpine AS builder

RUN apk update && apk add --no-cache git curl

# Install Air
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b /usr/local/bin

# Install go-migrate
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Runtime stage
FROM golang:1.24-alpine

RUN apk update && apk add --no-cache curl postgresql-client

# Copy Air and migrate binaries
COPY --from=builder /usr/local/bin/air /usr/local/bin/air
COPY --from=builder /go/bin/migrate /usr/local/bin/migrate

WORKDIR /app
COPY . .

# Make entrypoint.sh executable
RUN chmod +x entrypoint.sh

# Non-root user
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser

# Healthcheck
HEALTHCHECK --interval=30s --timeout=3s --retries=3 \
  CMD curl -f http://localhost:8080/v1/health || exit 1

ENTRYPOINT ["./entrypoint.sh"]
