#!/bin/sh
set -e
until pg_isready -h "$POSTGRES_HOST" -p "$POSTGRES_PORT" -U "$POSTGRES_USER" -d "$POSTGRES_DB"; do
  echo "Waiting for PostgreSQL to be ready..."
  sleep 2
done
echo "POSTGRES_URL: $POSTGRES_URL"
echo "Applying database migrations..."
migrate -path /app/migrations -database "$POSTGRES_URL" -verbose up
echo "Starting application..."
exec air
