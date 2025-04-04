-- name: GetAuthor :one
SELECT * FROM authors
WHERE id = $1;

-- name: GetAuthorByName :one
SELECT * FROM authors
WHERE name = $1;

-- name: ListAuthors :many
SELECT * FROM authors
ORDER BY name
LIMIT $1 OFFSET $2;

-- name: CreateAuthor :one
INSERT INTO authors (
  name, bio, openlibrary_key, photos, alternate_names, personal_name, links, birth_date, death_date
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9
)
RETURNING *;

-- name: UpdateAuthor :one
UPDATE authors
SET 
  name = $2,
  bio = $3,
  openlibrary_key = $4,
  photos = $5,
  alternate_names = $6,
  personal_name = $7,
  links = $8,
  birth_date = $9,
  death_date = $10,
  updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;

-- name: DeleteAuthor :exec
DELETE FROM authors
WHERE id = $1;

-- name: ListAuthorsByBookID :many
SELECT a.* FROM authors a
JOIN book_authors ba ON a.id = ba.author_id
WHERE ba.book_id = $1
ORDER BY a.name;

-- name: AddBookAuthor :exec
INSERT INTO book_authors (
  book_id, author_id
) VALUES (
  $1, $2
);

-- name: RemoveBookAuthor :exec
DELETE FROM book_authors
WHERE book_id = $1 AND author_id = $2;

-- name: RemoveAllBookAuthors :exec
DELETE FROM book_authors
WHERE book_id = $1;
