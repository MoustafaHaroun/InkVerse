package repository

import (
	"database/sql"

	"github.com/google/uuid"
)

type Novel struct {
	ID        uuid.UUID `json:"id"`
	AuthorID  uuid.UUID `json:"author_id"`
	Title     string    `json:"title"`
	Synopsis  string    `json:"synopsis"`
	Rating    float64   `json:"rating"`
	CreatedAt string    `json:"created_at"`
}

type NovelRepository interface {
	GetAllNovels() ([]Novel, error)
}

type SQLNovelRepository struct {
	DB *sql.DB
}

func (r *SQLNovelRepository) GetAllNovels() ([]Novel, error) {
	query := "SELECT id, author_id, title, synopsis, rating, created_at FROM novels"
	rows, err := r.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var novels []Novel

	for rows.Next() {
		var novel Novel

		if err := rows.Scan(&novel.ID, &novel.AuthorID, &novel.Title, &novel.Synopsis, &novel.Rating, &novel.CreatedAt); err != nil {
			return nil, err
		}

		novels = append(novels, novel)
	}

	return novels, nil
}
