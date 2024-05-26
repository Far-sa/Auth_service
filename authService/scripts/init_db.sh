#!/bin/bash

set -e

# Function to check if PostgreSQL is ready
function wait_for_postgres() {
  until pg_isready -h postgres -U "$POSTGRES_USER"; do
    echo "Waiting for PostgreSQL..."
    sleep 2
  done
}

# Function to create the database if it doesn't exist
function create_database() {
  echo "Checking if database $AUTH_SERVICE_DB_NAME exists..."
  if psql -lqt -h postgres -U "$POSTGRES_USER" | cut -d \| -f 1 | grep -qw "$AUTH_SERVICE_DB_NAME"; then
    echo "Database $AUTH_SERVICE_DB_NAME already exists."
  else
    echo "Creating database $AUTH_SERVICE_DB_NAME..."
    createdb -h postgres -U "$POSTGRES_USER" "$AUTH_SERVICE_DB_NAME"
  fi
}

# Run the functions
wait_for_postgres
create_database

# Run migrations
echo "Running migrations..."
migrate -path /migrations -database "postgresql://$POSTGRES_USER:$POSTGRES_PASSWORD@postgres:5432/$AUTH_SERVICE_DB_NAME?sslmode=disable" -verbose up

echo "Database initialization complete."


# #!/bin/bash
# set -e

# # Load environment variables
# if [ -f /app/.env ]; then
#   export $(cat /app/.env | xargs)
# fi

# # Create role for auth-service
# psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" <<-EOSQL
#   CREATE ROLE ${AUTH_SERVICE_DB_USER} WITH LOGIN PASSWORD '${AUTH_SERVICE_DB_PASSWORD}';
#   CREATE DATABASE ${AUTH_SERVICE_DB_NAME};
#   GRANT ALL PRIVILEGES ON DATABASE ${AUTH_SERVICE_DB_NAME} TO ${AUTH_SERVICE_DB_USER};
# EOSQL

# # Create role for authz-service
# psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" <<-EOSQL
#   CREATE ROLE ${AUTHZ_SERVICE_DB_USER} WITH LOGIN PASSWORD '${AUTHZ_SERVICE_DB_PASSWORD}';
#   CREATE DATABASE ${AUTHZ_SERVICE_DB_NAME};
#   GRANT ALL PRIVILEGES ON DATABASE ${AUTHZ_SERVICE_DB_NAME} TO ${AUTHZ_SERVICE_DB_USER};
# EOSQL



# #!/bin/bash
# set -e

# # Load environment variables from .env file
# set -a
# source /docker-entrypoint-initdb.d/.env
# set +a

# # Define PostgreSQL host and port
# POSTGRES_HOST=postgres-auth
# POSTGRES_PORT=5432

# # Wait for PostgreSQL to start
# until pg_isready -h "$POSTGRES_HOST" -p "$POSTGRES_PORT" -U "$POSTGRES_USER"; do
#   >&2 echo "Postgres is unavailable - sleeping"
#   sleep 1
# done

# # Create a new role and user for the auth service
# psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$AUTH_SERVICE_DB_NAME" <<-EOSQL
#   CREATE ROLE $AUTH_SERVICE_DB_USER WITH LOGIN PASSWORD '$AUTH_SERVICE_DB_PASSWORD';
#   GRANT ALL PRIVILEGES ON DATABASE $AUTH_SERVICE_DB_NAME TO $AUTH_SERVICE_DB_USER;
# EOSQL

# echo "Role and user for auth service created successfully."

