package chapter

import (
	"github.com/google/uuid"
)

type ChapterService struct {
	ChapterRepository ChapterRepository
}

func (s *ChapterService) GetByNovelId(novel_id uuid.UUID) ([]Chapter, error) {
	return s.ChapterRepository.GetByNovelId(novel_id)
}
