-- name: GetReview :one
SELECT * FROM book_reviews
WHERE id = $1;

-- name: GetReviewByUserAndBook :one
SELECT * FROM book_reviews
WHERE user_id = $1 AND book_id = $2;

-- name: ListReviews :many
SELECT * FROM book_reviews
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: ListReviewsByBookID :many
SELECT * FROM book_reviews
WHERE book_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: ListReviewsByUserID :many
SELECT * FROM book_reviews
WHERE user_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: CreateReview :one
INSERT INTO book_reviews (
  book_id, user_id, rating, review_text
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: UpdateReview :one
UPDATE book_reviews
SET 
  rating = $2,
  review_text = $3
WHERE id = $1
RETURNING *;

-- name: DeleteReview :exec
DELETE FROM book_reviews
WHERE id = $1;
