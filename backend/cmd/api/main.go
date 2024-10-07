package main

import (
	"errors"
	"log/slog"
	"net/http"
	"os"

	"github.com/MoustafaHaroun/InkVerse/internal/database"
	"github.com/MoustafaHaroun/InkVerse/internal/server/chapter"
	"github.com/MoustafaHaroun/InkVerse/internal/server/novel"
	"github.com/MoustafaHaroun/InkVerse/pkg/middleware"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	router := http.NewServeMux()

	// Add middleware stack
	stack := middleware.CreateStack(
		middleware.Logging,
		middleware.CORS,
	)

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
	novelRepository := &novel.SQLNovelRepository{DB: dbConn}
	novelService := novel.NovelService{NovelRepository: novelRepository}
	novelHandler := novel.NovelHandler{NovelService: novelService}

	router.HandleFunc("GET /novels/", novelHandler.GetAllNovelsHandler)
	router.HandleFunc("POST /novels/", novelHandler.AddNovelHandler)

	// chapter
	chapterRepository := &chapter.SQLChapterRepository{DB: dbConn}
	chapterService := chapter.ChapterService{ChapterRepository: chapterRepository}
	chapterHandler := chapter.ChapterHandler{ChapterService: chapterService}

	router.HandleFunc("GET /chapters/{id}", chapterHandler.GetByNovelIdHandler)
	router.HandleFunc("POST /chapters/", chapterHandler.AddChapterHandler)

	server := http.Server{
		Addr:    ":8000",
		Handler: stack(router),
	}

	slog.Info("Server starting", slog.String("addr", server.Addr))

	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("Failed to start server", slog.Any("error", err))
	}
}
