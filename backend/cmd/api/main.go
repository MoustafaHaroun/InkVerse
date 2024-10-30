package main

import (
	"errors"
	"log/slog"
	"net/http"
	"os"

	"github.com/MoustafaHaroun/InkVerse/internal/database"
	"github.com/MoustafaHaroun/InkVerse/internal/server/chapter"
	"github.com/MoustafaHaroun/InkVerse/internal/server/novel"
	"github.com/MoustafaHaroun/InkVerse/internal/server/user"
	"github.com/MoustafaHaroun/InkVerse/pkg/middleware"
)

func main() {
	setLogger()

	router := http.NewServeMux()

	// Add middleware stack
	stack := middleware.CreateStack(
		middleware.Logging,
		middleware.CORS,
	)

	// Open database connection
	dbConn := database.Connect()
	defer dbConn.Close()

	// user
	userRepository := user.NewUserRepository(dbConn)
	userHandler := user.NewUserHandler(*userRepository)
	userHandler.RegisterRoutes(router)

	// novel
	novelRepository := novel.NewSQLNovelRepository(dbConn)
	novelHandler := novel.NewNovelHandler(novelRepository)
	novelHandler.RegisterRoutes(router)

	// chapter
	chapterRepository := chapter.NewSQLChapterRepository(dbConn)
	chapterHandler := chapter.NewChapterHandler(chapterRepository, userRepository)
	chapterHandler.RegisterRoutes(router)

	server := http.Server{
		Addr:    ":8000",
		Handler: stack(router),
	}

	slog.Info("Server starting", slog.String("addr", server.Addr))

	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("Failed to start server", slog.Any("error", err))
	}
}

func setLogger() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
}
