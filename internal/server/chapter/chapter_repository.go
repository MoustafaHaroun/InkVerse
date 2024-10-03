package chapter

import (
	"database/sql"

	"github.com/google/uuid"
)

type Chapter struct {
	ID      uuid.UUID `json:"id"`
	NovelID uuid.UUID `json:"novel_id"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
}

type ChapterRepository interface {
	GetByNovelId(novel_id uuid.UUID) ([]Chapter, error)
	AddChapter(novel_id uuid.UUID, title string, content string) error
}

type SQLChapterRepository struct {
	DB *sql.DB
}

func (r *SQLChapterRepository) GetByNovelId(novel_id uuid.UUID) ([]Chapter, error) {
	query := "SELECT chapter_id, novel_id, title, content FROM chapters WHERE novel_id = $1"

	rows, err := r.DB.Query(query, novel_id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var chapters []Chapter

	for rows.Next() {
		var chapter Chapter

		if err := rows.Scan(&chapter.ID, &chapter.NovelID, &chapter.Title, &chapter.Content); err != nil {
			return nil, err
		}

		chapters = append(chapters, chapter)
	}

	return chapters, nil
}

func (r *SQLChapterRepository) AddChapter(novel_id uuid.UUID, title string, content string) error {
	query := "INSERT INTO chapters (novel_id, title, content) VALUES ($1,$2,$3)"

	_, err := r.DB.Query(query, novel_id, title, content)

	if err != nil {
		return err
	}

	return err
}
