-- name: GetBook :one
SELECT * FROM books
WHERE id = $1;

-- name: ListBooks :many
SELECT * FROM books
ORDER BY title
LIMIT $1 OFFSET $2;

-- name: SearchBooks :many
SELECT * FROM books
WHERE 
  title ILIKE '%' || $1 || '%'
  OR publisher ILIKE '%' || $1 || '%'
  OR description ILIKE '%' || $1 || '%'
ORDER BY title
LIMIT $2 OFFSET $3;

-- name: CreateBook :one
INSERT INTO books (
  isbn_10, isbn_13, title, publisher,
  published_date, description, page_count, language,
  thumbnail_url, total_copies, available_copies
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
)
RETURNING *;

-- name: UpdateBook :one
UPDATE books
SET 
  isbn_10 = $2,
  isbn_13 = $3,
  title = $4,
  publisher = $5,
  published_date = $6,
  description = $7,
  page_count = $8,
  language = $9,
  thumbnail_url = $10,
  total_copies = $11,
  available_copies = $12,
  updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;

-- name: DeleteBook :exec
DELETE FROM books
WHERE id = $1;

-- name: UpdateBookCopies :one
UPDATE books
SET 
  total_copies = $2,
  available_copies = $3,
  updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;
