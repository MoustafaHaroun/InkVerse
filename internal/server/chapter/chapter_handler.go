package chapter

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
)

type ChapterHandler struct {
	ChapterService ChapterService
}

func (h *ChapterHandler) GetByNovelIdHandler(w http.ResponseWriter, r *http.Request) {
	novel_id := r.PathValue("id")

	novels, err := h.ChapterService.GetByNovelId(uuid.MustParse(novel_id))

	if novels == nil {
		http.Error(w, "No chapters found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Failed to get chapters", http.StatusInternalServerError)
		slog.Error("Failed to get chapters", slog.Any("err", err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(novels)
}

func (h *ChapterHandler) AddChapterHandler(w http.ResponseWriter, r *http.Request) {
	var chapter struct {
		NovelID uuid.UUID `json:"novel_id"`
		Title   string    `json:"title"`
		Content string    `json:"content"`
	}

	if err := json.NewDecoder(r.Body).Decode(&chapter); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := h.ChapterService.AddChapter(chapter.NovelID, chapter.Title, chapter.Content); err != nil {
		http.Error(w, "Failed to add chapter", http.StatusInternalServerError)
		slog.Error("Failed to add chapter:", slog.Any("error", err))
		return
	}

	w.WriteHeader(http.StatusCreated)
}
