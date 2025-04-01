-- name: GetCategory :one
SELECT * FROM categories
WHERE id = $1;

-- name: GetCategoryByName :one
SELECT * FROM categories
WHERE name = $1;

-- name: ListCategories :many
SELECT * FROM categories
ORDER BY name
LIMIT $1 OFFSET $2;

-- name: CreateCategory :one
INSERT INTO categories (name)
VALUES ($1)
RETURNING *;

-- name: UpdateCategory :one
UPDATE categories
SET 
  name = $2,
  updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;

-- name: DeleteCategory :exec
DELETE FROM categories
WHERE id = $1;

-- name: ListCategoriesByBookID :many
SELECT c.* FROM categories c
JOIN book_categories bc ON c.id = bc.category_id
WHERE bc.book_id = $1
ORDER BY c.name;

-- name: AddBookCategory :exec
INSERT INTO book_categories (
  book_id, category_id
) VALUES (
  $1, $2
);

-- name: RemoveBookCategory :exec
DELETE FROM book_categories
WHERE book_id = $1 AND category_id = $2;

-- name: RemoveAllBookCategories :exec
DELETE FROM book_categories
WHERE book_id = $1;
