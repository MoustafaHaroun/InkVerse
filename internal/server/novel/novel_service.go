package novel

import (
	"errors"
	"strings"

	"github.com/google/uuid"
)

type NovelService struct {
	NovelRepository NovelRepository
}

const MaxSynopsisLen int = 800

var ErrSynopsisLenExceeded = errors.New("a novel synopsis can only be 800 words long")

func (s *NovelService) GetAllNovels() ([]Novel, error) {
	return s.NovelRepository.GetAllNovels()
}

func (s *NovelService) AddNovel(author_id uuid.UUID, title string, synopsis string) error {
	if len(strings.Split(synopsis, "")) > MaxSynopsisLen {
		return ErrSynopsisLenExceeded
	}

	return s.NovelRepository.AddNovel(author_id, title, synopsis)
}
