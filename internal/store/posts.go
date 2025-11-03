package store

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

type PostStore struct {
	db *sql.DB
}

type Post struct {
	ID      int64  `db:"id" json:"id"`
	Content string `db:"content" json:"content"`

	Title string `db:"title" json:"title"`

	AuthorID int64 `db:"author_id" json:"author_id"`

	Tags []string `db:"tags" json:"tags"`

	CreatedAt string `db:"created_at" json:"created_at"`

	UpdatedAt string `db:"updated_at" json:"updated_at"`
}

func (s *PostStore) Create(ctx context.Context, post *Post) error {
	query := `
	INSERT INTO posts (content, title, author_id, tags)
	VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at
	`

	err := s.db.QueryRowContext(ctx, query,
		post.Content,
		post.Title,
		post.AuthorID,
		pq.Array(post.Tags),
	).Scan(
		&post.ID,
		&post.CreatedAt,
		&post.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}
