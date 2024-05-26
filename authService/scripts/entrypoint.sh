#!/bin/sh

# Wait for the database to be ready
until pg_isready -h postgres -U root; do
  echo "Waiting for PostgreSQL..."
  sleep 2
done

# Create the database if it doesn't exist
psql -h postgres -U root -tc "SELECT 1 FROM pg_database WHERE datname = 'auth-db'" | grep -q 1 || psql -h postgres -U root -c "CREATE DATABASE \"auth-db\""

# Run database migrations
migrate -path /app/infrastructure/database/migrations -database "$DATABASE_URL" -verbose up

# Start the auth service
/app/auth-service
