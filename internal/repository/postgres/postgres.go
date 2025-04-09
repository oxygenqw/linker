package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func NewPostgresConnection(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open connection: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping connection: %w", err)
	}

	return db, nil
}
