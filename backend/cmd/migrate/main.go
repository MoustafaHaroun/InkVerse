package main

import (
	"log"

	"github.com/MoustafaHaroun/InkVerse/internal/database"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
)

func main() {
	dbConn := database.Connect()
	defer dbConn.Close()

	driver, err := postgres.WithInstance(dbConn, &postgres.Config{})
	if err != nil {
		log.Fatalf("Could not start SQL driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://cmd/migrate/migrations", "postgres", driver)
	if err != nil {
		log.Fatalf("Could not start migration: %v", err)
	}

	//TODO: maybe install there cli tools
	// if cmd == "down" {
	if err := m.Down(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Migration failed down: %v", err)
	}
	// }

	// cmd := os.Args[(len(os.Args) - 1)]
	// if cmd == "up" {
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Migration failed up: %v", err)
	}
	// }

	log.Printf("Migration has successfully been done!")
}
