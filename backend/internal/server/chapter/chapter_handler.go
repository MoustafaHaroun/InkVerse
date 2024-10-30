package chapter

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/MoustafaHaroun/InkVerse/internal/server/user"
	"github.com/MoustafaHaroun/InkVerse/pkg/middleware"
	"github.com/MoustafaHaroun/InkVerse/pkg/util"
	"github.com/google/uuid"
)

type ChapterHandler struct {
	ChapterRepository ChapterRepository
	UserRepository    user.UserRepository
}

func NewChapterHandler(chapterRepository ChapterRepository, userRepository user.UserRepository) *ChapterHandler {
	return &ChapterHandler{
		ChapterRepository: chapterRepository,
		UserRepository:    userRepository,
	}
}

func (h *ChapterHandler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /chapters/{id}", h.GetByIdHandler)
	router.HandleFunc("POST /chapters/", middleware.WithJWTAuth(h.AddChapterHandler, h.UserRepository))
	router.HandleFunc("GET /novels/{id}/chapters", h.GetByNovelIdHandler) //TODO: change this maybe give the chapter handler then chapter service
}

type ChapterDto struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Created_at  string    `json:"created_at"`
	Modified_at string    `json:"modified_at"`
}

func (h *ChapterHandler) GetByNovelIdHandler(w http.ResponseWriter, r *http.Request) {
	// Parse id to uuid
	novel_id, err := uuid.Parse(r.PathValue("id"))

	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Get novel
	chapters, err := h.ChapterRepository.GetByNovelId(novel_id)

	if err != nil {
		slog.Error("Failed to get chapters", slog.Any("err", err))
		util.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if len(chapters) == 0 {
		util.WriteError(w, http.StatusNotFound, nil)
		return
	}

	//TODO: refactor this later to a smarter method.
	chapterDTOs := make([]ChapterDto, len(chapters))
	for i, chapter := range chapters {
		chapterDTOs[i] = ChapterDto{
			ID:          chapter.ID,
			Title:       chapter.Title,
			Created_at:  chapter.CreatedAt,
			Modified_at: chapter.ModifiedAt,
		}
	}

	util.WriteJSON(w, http.StatusOK, chapterDTOs)
}

func (h *ChapterHandler) GetByIdHandler(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))

	if err != nil {
		http.Error(w, "invalid id: "+id.String(), http.StatusBadRequest)
		return
	}

	chapter, err := h.ChapterRepository.GetById(id)

	if err != nil {
		http.Error(w, "Failed to get chapter", http.StatusInternalServerError)
		slog.Error("Failed to get chapter", slog.Any("err", err))
		return
	}

	if chapter == nil {
		http.Error(w, "Could not find the novel with id: "+id.String(), http.StatusNotFound)
		return
	}

	util.WriteJSON(w, http.StatusOK, chapter)
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

	if err := h.ChapterRepository.Add(chapter.NovelID, chapter.Title, chapter.Content); err != nil {
		http.Error(w, "Failed to add chapter", http.StatusInternalServerError)
		slog.Error("Failed to add chapter:", slog.Any("error", err))
		return
	}

	util.WriteJSON(w, http.StatusCreated, nil)
}
