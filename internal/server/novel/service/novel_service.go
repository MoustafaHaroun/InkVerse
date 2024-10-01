package service

import (
	"github.com/MoustafaHaroun/InkVerse/internal/server/novel/repository"
	"github.com/google/uuid"
)

type NovelService struct {
	NovelRepository repository.NovelRepository
}

func (s *NovelService) GetAllNovels() ([]repository.Novel, error) {
	return s.NovelRepository.GetAllNovels()
}

func (s *NovelService) AddNovel(author_id uuid.UUID, title string, synopsis string) error {
	return s.NovelRepository.AddNovel(author_id, title, synopsis)
}
