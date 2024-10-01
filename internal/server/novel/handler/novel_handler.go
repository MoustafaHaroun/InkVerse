package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/MoustafaHaroun/InkVerse/internal/server/novel/service"
)

type NovelHandler struct {
	NovelService service.NovelService
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
