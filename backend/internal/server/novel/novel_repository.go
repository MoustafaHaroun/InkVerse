package novel

import (
	"database/sql"
	"errors"

	"github.com/MoustafaHaroun/InkVerse/internal/database"
	"github.com/google/uuid"
)

type Novel struct {
	ID         uuid.UUID `json:"id"`
	AuthorID   uuid.UUID `json:"author_id"`
	Title      string    `json:"title"`
	Synopsis   string    `json:"synopsis"`
	Rating     float64   `json:"rating"`
	CreatedAt  string    `json:"created_at"`
	ModifiedAt string    `json:"modified_at"`
}

var (
	ErrNovelExistsWithTitle = errors.New("novel with this title already exists")
)

type NovelRepository interface {
	// GetById retrieves a novel by its ID.
	GetById(id uuid.UUID) (*Novel, error)

	// GetAll retrieves all novels from the data storage.
	GetAll() ([]Novel, error)

	// Add adds a new novel.
	Add(author_id uuid.UUID, title string, synopsis string) error
}

type SQLNovelRepository struct {
	DB *sql.DB
}

func NewSQLNovelRepository(db *sql.DB) *SQLNovelRepository {
	return &SQLNovelRepository{
		DB: db,
	}
}

func (r *SQLNovelRepository) GetAll() ([]Novel, error) {
	query := "SELECT novel_id, author_id, title, synopsis, rating, created_at, modified_at FROM novels"
	rows, err := r.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var novels []Novel

	for rows.Next() {
		var novel Novel

		if err := rows.Scan(&novel.ID, &novel.AuthorID, &novel.Title, &novel.Synopsis, &novel.Rating, &novel.CreatedAt, &novel.ModifiedAt); err != nil {
			return nil, err
		}

		novels = append(novels, novel)
	}

	return novels, nil
}

func (r *SQLNovelRepository) GetById(id uuid.UUID) (*Novel, error) {
	query := "SELECT novel_id, author_id, title, synopsis, rating, created_at modified_at FROM novels where novel_id = $1"
	rows, err := r.DB.Query(query, id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var novel Novel

	for rows.Next() {
		if err := rows.Scan(&novel.ID, &novel.AuthorID, &novel.Title, &novel.Synopsis, &novel.Rating, &novel.CreatedAt, &novel.ModifiedAt); err != nil {
			return nil, err
		}
		return &novel, nil
	}

	return nil, sql.ErrNoRows
}

func (r *SQLNovelRepository) Add(author_id uuid.UUID, title string, synopsis string) error {
	query := "INSERT INTO novels (author_id, title, synopsis) VALUES ($1,$2,$3)"
	_, err := r.DB.Exec(query, author_id, title, synopsis)

	if err != nil {
		if database.IsUniqueViolation(err) {
			return ErrNovelExistsWithTitle
		}
	}
	return err
}
