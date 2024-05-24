migrateup:
	migrate -path infrastructure/database/migrations -database "postgresql://root:password@localhost:5432/PGU?sslmode=disable" -verbose up

migratedown:
	migrate -path infrastructure/database/migrations -database "postgresql://root:password@localhost:5432/PGU?sslmode=disable" -verbose down

# Commands
.PHONY: migrateup migratedown


#! sqlc generate
#! migrate create -ext sql -dir migrations create_users_table
#! migrate create -ext sql -dir infrastructure/database/migrations init_database
