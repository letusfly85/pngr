// Code generated by sqlc. DO NOT EDIT.
// source: post.sql

package db

import (
	"context"
	"time"
)

const createPost = `-- name: CreatePost :one
INSERT INTO posts (author_id, title, body, status) VALUES ($1, $2, $3, $4) RETURNING id, author_id, title, body, status, created_at, updated_at
`

type CreatePostParams struct {
	AuthorID int64      `json:"author_id"`
	Title    string     `json:"title"`
	Body     string     `json:"body"`
	Status   PostStatus `json:"status"`
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, createPost,
		arg.AuthorID,
		arg.Title,
		arg.Body,
		arg.Status,
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.AuthorID,
		&i.Title,
		&i.Body,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deletePostByIDs = `-- name: DeletePostByIDs :exec
DELETE FROM posts WHERE author_id = $1 AND id = $2
`

type DeletePostByIDsParams struct {
	AuthorID int64 `json:"author_id"`
	ID       int64 `json:"id"`
}

func (q *Queries) DeletePostByIDs(ctx context.Context, arg DeletePostByIDsParams) error {
	_, err := q.db.ExecContext(ctx, deletePostByIDs, arg.AuthorID, arg.ID)
	return err
}

const findPostByIDs = `-- name: FindPostByIDs :one
SELECT id, author_id, title, body, status, created_at, updated_at FROM posts WHERE author_id = $1 AND id = $2 LIMIT 1
`

type FindPostByIDsParams struct {
	AuthorID int64 `json:"author_id"`
	ID       int64 `json:"id"`
}

func (q *Queries) FindPostByIDs(ctx context.Context, arg FindPostByIDsParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, findPostByIDs, arg.AuthorID, arg.ID)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.AuthorID,
		&i.Title,
		&i.Body,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findPostsByAuthor = `-- name: FindPostsByAuthor :many
SELECT id, author_id, title, body, status, created_at, updated_at FROM posts WHERE author_id = $1 ORDER BY id DESC
`

func (q *Queries) FindPostsByAuthor(ctx context.Context, authorID int64) ([]Post, error) {
	rows, err := q.db.QueryContext(ctx, findPostsByAuthor, authorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Post{}
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.AuthorID,
			&i.Title,
			&i.Body,
			&i.Status,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updatePost = `-- name: UpdatePost :one
UPDATE posts SET title = $1, body = $2, updated_at = $3 WHERE id = $4 AND author_id = $5 RETURNING id, author_id, title, body, status, created_at, updated_at
`

type UpdatePostParams struct {
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	UpdatedAt time.Time `json:"updated_at"`
	ID        int64     `json:"id"`
	AuthorID  int64     `json:"author_id"`
}

func (q *Queries) UpdatePost(ctx context.Context, arg UpdatePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, updatePost,
		arg.Title,
		arg.Body,
		arg.UpdatedAt,
		arg.ID,
		arg.AuthorID,
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.AuthorID,
		&i.Title,
		&i.Body,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
