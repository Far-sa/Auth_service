package repository

import "user-service/infrastructure/database"

type DB struct {
	conn *database.SqlDB
}

func New(conn *database.SqlDB) *DB {
	return &DB{conn: conn}
}
