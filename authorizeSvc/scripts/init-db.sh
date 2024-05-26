#!/bin/bash
set -e

# Load environment variables
if [ -f /app/.env ]; then
  export $(cat /app/.env | xargs)
fi

# Create role for auth-service
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" <<-EOSQL
  CREATE ROLE ${AUTH_SERVICE_DB_USER} WITH LOGIN PASSWORD '${AUTH_SERVICE_DB_PASSWORD}';
  CREATE DATABASE ${AUTH_SERVICE_DB_NAME};
  GRANT ALL PRIVILEGES ON DATABASE ${AUTH_SERVICE_DB_NAME} TO ${AUTH_SERVICE_DB_USER};
EOSQL

# Create role for authz-service
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" <<-EOSQL
  CREATE ROLE ${AUTHZ_SERVICE_DB_USER} WITH LOGIN PASSWORD '${AUTHZ_SERVICE_DB_PASSWORD}';
  CREATE DATABASE ${AUTHZ_SERVICE_DB_NAME};
  GRANT ALL PRIVILEGES ON DATABASE ${AUTHZ_SERVICE_DB_NAME} TO ${AUTHZ_SERVICE_DB_USER};
EOSQL



# #!/bin/bash
# set -e

# # Source the shared .env file
# #source ../../../.env

# # Load environment variables from .env file
# set -a
# source /docker-entrypoint-initdb.d/.env
# set +a

# # Define PostgreSQL host and port
# POSTGRES_HOST=postgres-authz
# POSTGRES_PORT=5432

# # Wait for PostgreSQL to start
# until pg_isready -h "$POSTGRES_HOST" -p "$POSTGRES_PORT" -U "$POSTGRES_USER"; do
#   >&2 echo "Postgres is unavailable - sleeping"
#   sleep 1
# done

# # Create a new role and user for the authorization service
# psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$AUTHZ_SERVICE_DB_NAME" <<-EOSQL
#   CREATE ROLE $AUTHZ_SERVICE_DB_USER WITH LOGIN PASSWORD '$AUTHZ_SERVICE_DB_PASSWORD';
#   GRANT ALL PRIVILEGES ON DATABASE $AUTHZ_SERVICE_DB_NAME TO $AUTHZ_SERVICE_DB_USER;
# EOSQL

# echo "Role and user for authorization service created successfully."

