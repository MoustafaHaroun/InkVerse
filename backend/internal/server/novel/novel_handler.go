package novel

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/MoustafaHaroun/InkVerse/pkg/util"
	"github.com/google/uuid"
)

type NovelHandler struct {
	NovelRepository NovelRepository
}

func NewNovelHandler(novelRepository NovelRepository) *NovelHandler {
	return &NovelHandler{
		NovelRepository: novelRepository,
	}
}

func (h *NovelHandler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /novels/", h.GetAllNovels)
	// router.HandleFunc("GET /novels/{id}", h.GetNovelById)
	router.HandleFunc("POST /novels/", h.AddNovel)
}

// Get's all the novels from the database
func (h *NovelHandler) GetAllNovels(w http.ResponseWriter, r *http.Request) {
	novels, err := h.NovelRepository.GetAll()

	if err != nil {
		slog.Error("Failed to get novels: %v", slog.Any("err", err))
		util.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	util.WriteJSON(w, http.StatusOK, novels)
}

// Get a novel based on the ID
func (h *NovelHandler) GetNovelById(w http.ResponseWriter, r *http.Request) {
	novel_id, err := uuid.Parse(r.PathValue("id"))

	if err != nil {
		http.Error(w, "Invalid id:"+novel_id.String(), http.StatusBadRequest)
		return
	}

	novel, err := h.NovelRepository.GetById(novel_id)

	if err != nil {
		slog.Error("Failed to retrieve the novel", slog.Any("err:", err))
		util.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to get novel"))
		return
	}

	if novel == nil {
		http.Error(w, "Could not find the novel with id: "+novel_id.String(), http.StatusNotFound)
		return
	}

	util.WriteJSON(w, http.StatusOK, novel)
}

// Add a novel to the database
func (h *NovelHandler) AddNovel(w http.ResponseWriter, r *http.Request) {
	var novel struct {
		AuthorID uuid.UUID `json:"author_id"`
		Title    string    `json:"title"`
		Synopsis string    `json:"synopsis"`
	}

	if err := util.ParseJsonPayload(r, novel); err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
	}

	if err := h.NovelRepository.Add(novel.AuthorID, novel.Title, novel.Title); err != nil {
		slog.Error("Failed to add novel:", slog.Any("error", err))
		util.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	util.WriteJSON(w, http.StatusCreated, nil)
}
