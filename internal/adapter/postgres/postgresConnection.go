package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectToPostgresDB(host string, port int, user, password, dbname string) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to connect from PostgreSQL: %v", err)
	}
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping PostgreSQL: %v", err)
	}
	return db, nil
}

func DisconnectFromPostgresDB(db *sql.DB) error {
	err := db.Close()
	if err != nil {
		return fmt.Errorf("failed to disconnect from PostgreSQL: %v", err)
	}
	return nil
}
