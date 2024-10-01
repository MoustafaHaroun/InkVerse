package database

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file" // Import the file source driver
	_ "github.com/lib/pq"                                // PostgreSQL driver
)

func Connect() *sql.DB {
	connString := "postgresql://admin:admin@localhost:5432/app?sslmode=disable"

	db, err := sql.Open("postgres", connString)

	// Check if connection has been established with the database.
	if err != nil && db.Ping() != nil {
		log.Fatalf("Failed to make a connection to the database: %v", err)
	}

	log.Printf("Connected to database has been established!")

	return db
}

func Migrate(db *sql.DB) {
	log.Printf("Migration running...")

	// Get the driver of postgresql
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("Could not start SQL driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://internal/database/migrations", "postgres", driver)
	if err != nil {
		log.Fatalf("Could not start migration: %v", err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Printf("Migration has successfully been done!")
}
