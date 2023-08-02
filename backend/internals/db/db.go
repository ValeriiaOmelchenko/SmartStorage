package db

import (
	"database/sql"
	"time"
)

// PostgresDBRepo is the struct used to wrap our database connection pool, so that we
// can easily swap out a real database for a test database, or move to another database
// entirely, as long as the thing being swapped implements all of the functions in the type
// repository.DatabaseRepo.
type PostgresDB struct {
	DB *sql.DB
}

const dbTimeout = time.Second * 3

// Connection returns underlying connection pool.
func (m *PostgresDB) Connection() *sql.DB {
	return m.DB
}
