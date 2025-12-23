package database

import (
	"database/sql"
	"fmt"
	"gstat/internal/configuration"
	"gstat/internal/protocol"
	"time"

	_ "github.com/lib/pq"
)

const createResultHistoryTable = `
CREATE TABLE IF NOT EXISTS result_history (
    id BIGSERIAL PRIMARY KEY,

    target TEXT NOT NULL,
    protocol TEXT NOT NULL,

    reachable BOOLEAN NOT NULL,
    message TEXT,

    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
`

func Open(database configuration.Database) (*sql.DB, error) {
	connectionQuery := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		database.Host,
		database.Port,
		database.User,
		database.Password,
		database.Name,
		func() string {
			if database.Ssl {
				return "enable"
			}
			return "disable"
		}(),
	)

	db, err := sql.Open("postgres", connectionQuery)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func Close(db *sql.DB) error {
	return db.Close()
}

func EnsureTables(db *sql.DB) error {
	_, err := db.Exec(createResultHistoryTable)
	return err
}

func InsertResult(db *sql.DB, result protocol.Result, time time.Time) error {
	_, err := db.Exec(
		`INSERT INTO result_history (target, protocol, reachable, message, created_at) VALUES ($1, $2, $3, $4, $5)`,
		result.Target,
		string(result.Protocol),
		result.Reachable,
		result.Message,
		time,
	)
	return err
}
