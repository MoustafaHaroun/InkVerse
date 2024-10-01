package main

import (
	"errors"
	"log/slog"
	"net/http"
	"os"

	"github.com/MoustafaHaroun/InkVerse/internal/database"
	"github.com/MoustafaHaroun/InkVerse/internal/server/novel/handler"
	"github.com/MoustafaHaroun/InkVerse/internal/server/novel/repository"
	"github.com/MoustafaHaroun/InkVerse/internal/server/novel/service"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	router := http.NewServeMux()

	// Open database connection
	dbConn := database.Connect()
	defer dbConn.Close()

	// Run Migrations
	database.Migrate(dbConn)

	// Health check
	router.HandleFunc("/{$}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// novels
	novelRepository := &repository.SQLNovelRepository{DB: dbConn}
	novelService := service.NovelService{NovelRepository: novelRepository}
	novelHandler := handler.NovelHandler{NovelService: novelService}

	router.HandleFunc("GET /novels/", novelHandler.GetAllNovelsHandler)

	server := http.Server{
		Addr:    ":8000",
		Handler: router,
	}

	slog.Info("Server starting", slog.String("addr", server.Addr))

	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("Failed to start server", slog.Any("error", err))
	}
}
