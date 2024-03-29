package database

import (
	"sync"

	"github.com/jackc/pgx"
)

var DB *pgx.ConnPool
var once sync.Once

const maxConn = 8

func Connect() (conn *pgx.ConnPool) {
	once.Do(func() {
		connConfig := pgx.ConnConfig{
			User:     "ksu",
			Password: "pswd",
			Host:     "localhost",
			Port:     5432,
			Database: "avito_chat_db",
		}

		DB, _ = pgx.NewConnPool(pgx.ConnPoolConfig{
			ConnConfig:     connConfig,
			MaxConnections: maxConn,
		})
	})

	return DB
}
