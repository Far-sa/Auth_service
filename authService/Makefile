createdb:
	docker exec -it postgres13 createdb --username=root --owner=root auth-db

dropdb:
	docker exec -it postgres13 dropdb auth-db

migrateup:
	migrate -path infrastructure/database/migrations -database "postgresql://root:password@localhost:5432/auth-db?sslmode=disable" -verbose up

migratedown:
	migrate -path infrastructure/database/migrations -database "postgresql://root:password@localhost:5432/auth_db?sslmode=disable" -verbose down

# Commands
.PHONY: createdb dropdb migrateup migratedown


#! sqlc generate
#! migrate create -ext sql -dir migrations create_users_table
#! migrate create -ext sql -dir infrastructure/database/migrations init_database

#?
# .PHONY: all create-user-db migrate build deploy

# all: build deploy

# create-user-db:
#        # Use psql or a script to create user and database (refer to environment variables for credentials)
#        psql -h <your_postgres_host> -U ${POSTGRES_USER} -P ${POSTGRES_PASSWORD} -c "CREATE USER ${AUTH_SERVICE_DB_USER} WITH PASSWORD '${AUTH_SERVICE_DB_PASSWORD}'; CREATE DATABASE ${AUTH_SERVICE_DB_NAME} OWNER ${AUTH_SERVICE_DB_USER};"

# migrate: create-user-db
#        # Execute Flyway migrations, replace placeholders with actual values
#        flyway -url jdbc:postgresql://<your_postgres_host>/<your_database_name> -user ${FLYWAY_USER} -password ${FLYWAY_PASSWORD} migrate

# build:
#        # Build your service binary using your chosen language's compiler (e.g., go build)
#        go build -o ./bin/auth-service ./cmd/auth-service

# deploy: build
#        # Commands to deploy your service using Docker or Kubernetes (replace with your specific deployment commands)
#        docker build -t your_service_image .

# # Optional target for cleaning up (remove built binary)
# clean:
#        rm -rf ./bin/auth-service
