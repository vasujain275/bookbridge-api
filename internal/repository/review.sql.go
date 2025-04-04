// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: review.sql

package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createReview = `-- name: CreateReview :one
INSERT INTO book_reviews (
  book_id, user_id, rating, review_text
) VALUES (
  $1, $2, $3, $4
)
RETURNING id, book_id, user_id, rating, review_text, created_at
`

type CreateReviewParams struct {
	BookID     uuid.UUID   `json:"book_id"`
	UserID     uuid.UUID   `json:"user_id"`
	Rating     int32       `json:"rating"`
	ReviewText pgtype.Text `json:"review_text"`
}

func (q *Queries) CreateReview(ctx context.Context, arg CreateReviewParams) (BookReview, error) {
	row := q.db.QueryRow(ctx, createReview,
		arg.BookID,
		arg.UserID,
		arg.Rating,
		arg.ReviewText,
	)
	var i BookReview
	err := row.Scan(
		&i.ID,
		&i.BookID,
		&i.UserID,
		&i.Rating,
		&i.ReviewText,
		&i.CreatedAt,
	)
	return i, err
}

const deleteReview = `-- name: DeleteReview :exec
DELETE FROM book_reviews
WHERE id = $1
`

func (q *Queries) DeleteReview(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteReview, id)
	return err
}

const getReview = `-- name: GetReview :one
SELECT id, book_id, user_id, rating, review_text, created_at FROM book_reviews
WHERE id = $1
`

func (q *Queries) GetReview(ctx context.Context, id uuid.UUID) (BookReview, error) {
	row := q.db.QueryRow(ctx, getReview, id)
	var i BookReview
	err := row.Scan(
		&i.ID,
		&i.BookID,
		&i.UserID,
		&i.Rating,
		&i.ReviewText,
		&i.CreatedAt,
	)
	return i, err
}

const getReviewByUserAndBook = `-- name: GetReviewByUserAndBook :one
SELECT id, book_id, user_id, rating, review_text, created_at FROM book_reviews
WHERE user_id = $1 AND book_id = $2
`

type GetReviewByUserAndBookParams struct {
	UserID uuid.UUID `json:"user_id"`
	BookID uuid.UUID `json:"book_id"`
}

func (q *Queries) GetReviewByUserAndBook(ctx context.Context, arg GetReviewByUserAndBookParams) (BookReview, error) {
	row := q.db.QueryRow(ctx, getReviewByUserAndBook, arg.UserID, arg.BookID)
	var i BookReview
	err := row.Scan(
		&i.ID,
		&i.BookID,
		&i.UserID,
		&i.Rating,
		&i.ReviewText,
		&i.CreatedAt,
	)
	return i, err
}

const listReviews = `-- name: ListReviews :many
SELECT id, book_id, user_id, rating, review_text, created_at FROM book_reviews
ORDER BY created_at DESC
LIMIT $1 OFFSET $2
`

type ListReviewsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListReviews(ctx context.Context, arg ListReviewsParams) ([]BookReview, error) {
	rows, err := q.db.Query(ctx, listReviews, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []BookReview
	for rows.Next() {
		var i BookReview
		if err := rows.Scan(
			&i.ID,
			&i.BookID,
			&i.UserID,
			&i.Rating,
			&i.ReviewText,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listReviewsByBookID = `-- name: ListReviewsByBookID :many
SELECT id, book_id, user_id, rating, review_text, created_at FROM book_reviews
WHERE book_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3
`

type ListReviewsByBookIDParams struct {
	BookID uuid.UUID `json:"book_id"`
	Limit  int32     `json:"limit"`
	Offset int32     `json:"offset"`
}

func (q *Queries) ListReviewsByBookID(ctx context.Context, arg ListReviewsByBookIDParams) ([]BookReview, error) {
	rows, err := q.db.Query(ctx, listReviewsByBookID, arg.BookID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []BookReview
	for rows.Next() {
		var i BookReview
		if err := rows.Scan(
			&i.ID,
			&i.BookID,
			&i.UserID,
			&i.Rating,
			&i.ReviewText,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listReviewsByUserID = `-- name: ListReviewsByUserID :many
SELECT id, book_id, user_id, rating, review_text, created_at FROM book_reviews
WHERE user_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3
`

type ListReviewsByUserIDParams struct {
	UserID uuid.UUID `json:"user_id"`
	Limit  int32     `json:"limit"`
	Offset int32     `json:"offset"`
}

func (q *Queries) ListReviewsByUserID(ctx context.Context, arg ListReviewsByUserIDParams) ([]BookReview, error) {
	rows, err := q.db.Query(ctx, listReviewsByUserID, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []BookReview
	for rows.Next() {
		var i BookReview
		if err := rows.Scan(
			&i.ID,
			&i.BookID,
			&i.UserID,
			&i.Rating,
			&i.ReviewText,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateReview = `-- name: UpdateReview :one
UPDATE book_reviews
SET 
  rating = $2,
  review_text = $3
WHERE id = $1
RETURNING id, book_id, user_id, rating, review_text, created_at
`

type UpdateReviewParams struct {
	ID         uuid.UUID   `json:"id"`
	Rating     int32       `json:"rating"`
	ReviewText pgtype.Text `json:"review_text"`
}

func (q *Queries) UpdateReview(ctx context.Context, arg UpdateReviewParams) (BookReview, error) {
	row := q.db.QueryRow(ctx, updateReview, arg.ID, arg.Rating, arg.ReviewText)
	var i BookReview
	err := row.Scan(
		&i.ID,
		&i.BookID,
		&i.UserID,
		&i.Rating,
		&i.ReviewText,
		&i.CreatedAt,
	)
	return i, err
}
