package service

import "github.com/MoustafaHaroun/InkVerse/internal/server/novel/repository"

type NovelService struct {
	NovelRepository repository.NovelRepository
}

func (s *NovelService) GetAllNovels() ([]repository.Novel, error) {
	return s.NovelRepository.GetAllNovels()
}
