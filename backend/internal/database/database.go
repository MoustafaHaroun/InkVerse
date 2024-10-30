package database

import (
	"database/sql"
	"log"
	"log/slog"

	_ "github.com/golang-migrate/migrate/v4/source/file" // Import the file source driver
	"github.com/lib/pq"
	_ "github.com/lib/pq" // PostgreSQL driver
)

func Connect() *sql.DB {
	connString := "postgresql://admin:admin@localhost:5432/app?sslmode=disable" //TODO: create this string based on a env

	db, err := sql.Open("postgres", connString)

	// Check if connection has been established with the database.
	if err != nil && db.Ping() != nil {
		log.Fatalf("Failed to make a connection to the database: %v", err)
	}

	slog.Info("Connected to database has been established!")

	return db
}

func IsUniqueViolation(err error) bool {
	if pqErr, ok := err.(*pq.Error); ok {
		if pqErr.Code == "23505" {
			return true
		}
	}
	return false
}
