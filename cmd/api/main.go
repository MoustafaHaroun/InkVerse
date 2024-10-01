package main

import (
	"errors"
	"log/slog"
	"net/http"
	"os"

	"github.com/MoustafaHaroun/InkVerse/internal/database"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	router := http.NewServeMux()

	// Open database connection
	dbConn := database.Connect()
	defer dbConn.Close()

	// Run Migrations
	database.Migrate(dbConn)

	server := http.Server{
		Addr:    ":8000",
		Handler: router,
	}

	logger.Info("Server starting", slog.String("addr", server.Addr))

	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		logger.Error("Failed to start server", slog.Any("error", err))
	}
}
