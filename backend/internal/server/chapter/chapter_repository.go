package chapter

import (
	"database/sql"

	"github.com/google/uuid"
)

type Chapter struct {
	ID         uuid.UUID `json:"id"`
	NovelID    uuid.UUID `json:"novel_id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	CreatedAt  string    `json:"created_at"`
	ModifiedAt string    `json:"modified_at"`
}

type ChapterRepository interface {
	GetById(id uuid.UUID) (*Chapter, error)
	GetByNovelId(novel_id uuid.UUID) ([]Chapter, error)
	Add(novel_id uuid.UUID, title string, content string) error
}

type SQLChapterRepository struct {
	DB *sql.DB
}

func NewSQLChapterRepository(db *sql.DB) *SQLChapterRepository {
	return &SQLChapterRepository{
		DB: db,
	}
}

func (r *SQLChapterRepository) GetByNovelId(novel_id uuid.UUID) ([]Chapter, error) {
	query := "SELECT chapter_id, title, created_at, modified_at FROM chapters WHERE novel_id = $1"

	rows, err := r.DB.Query(query, novel_id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var chapters []Chapter

	for rows.Next() {
		var chapter Chapter

		if err := rows.Scan(&chapter.ID, &chapter.Title, &chapter.CreatedAt, &chapter.ModifiedAt); err != nil {
			return nil, err
		}

		chapters = append(chapters, chapter)
	}

	return chapters, nil
}

func (r *SQLChapterRepository) GetById(id uuid.UUID) (*Chapter, error) {
	query := "select chapter_id, title, content, created_at, modified_at FROM chapters WHERE novel_id = $1"

	rows, err := r.DB.Query(query, id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var chapter Chapter

	for rows.Next() {
		if err := rows.Scan(&chapter.ID, &chapter.Title, &chapter.Content, &chapter.CreatedAt, &chapter.ModifiedAt); err != nil {
			return nil, err
		}

		return &chapter, nil
	}

	return nil, sql.ErrNoRows
}

func (r *SQLChapterRepository) Add(novel_id uuid.UUID, title string, content string) error {
	query := "INSERT INTO chapters (novel_id, title, content) VALUES ($1,$2,$3)"

	_, err := r.DB.Query(query, novel_id, title, content)

	if err != nil {
		return err
	}

	return err
}
