package novel

import (
	"encoding/json"
	"log"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
)

type NovelHandler struct {
	NovelService NovelService
}

func (h *NovelHandler) GetAllNovelsHandler(w http.ResponseWriter, r *http.Request) {
	novels, err := h.NovelService.GetAllNovels()

	if err != nil {
		http.Error(w, "Failed to get all the novels", http.StatusInternalServerError)
		log.Printf("Failed to get novels: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(novels)
}

func (h *NovelHandler) AddNovelHandler(w http.ResponseWriter, r *http.Request) {
	var novel struct {
		AuthorID uuid.UUID `json:"author_id"`
		Title    string    `json:"title"`
		Synopsis string    `json:"synopsis"`
	}

	if err := json.NewDecoder(r.Body).Decode(&novel); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	//TODO: Add more validation

	if err := h.NovelService.AddNovel(novel.AuthorID, novel.Title, novel.Title); err != nil {
		http.Error(w, "Failed to add novel", http.StatusInternalServerError)
		slog.Error("Failed to add novel:", slog.Any("error", err))
		return
	}

	w.WriteHeader(http.StatusCreated)
}
