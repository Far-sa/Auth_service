createdb:
	docker exec -it postgres12 createdb --username=root --owner=root authz-db

dropdb:
	docker exec -it postgres12 dropdb authz-db

migrateup:
	migrate -path infrastructure/database/migrations -database "postgresql://root:password@localhost:5432/authz-db?sslmode=disable" -verbose up

migratedown:
	migrate -path infrastructure/database/migrations -database "postgresql://root:password@localhost:5432/authz-db?sslmode=disable" -verbose down

# Commands
.PHONY: createdb dropdb migrateup migratedown


#! sqlc generate
#! migrate create -ext sql -dir migrations create_users_table
#! migrate create -ext sql -dir infrastructure/database/migrations init_database
