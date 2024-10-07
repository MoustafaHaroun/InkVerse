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

func (s *ChapterService) AddChapter(novel_id uuid.UUID, title string, content string) error {
	return s.ChapterRepository.AddChapter(novel_id, title, content)
}
