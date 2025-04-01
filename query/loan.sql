-- name: GetLoan :one
SELECT * FROM loans
WHERE id = $1;

-- name: ListLoans :many
SELECT * FROM loans
ORDER BY borrowed_date DESC
LIMIT $1 OFFSET $2;

-- name: ListLoansByUserID :many
SELECT * FROM loans
WHERE user_id = $1
ORDER BY borrowed_date DESC
LIMIT $2 OFFSET $3;

-- name: ListLoansByBookID :many
SELECT * FROM loans
WHERE book_id = $1
ORDER BY borrowed_date DESC
LIMIT $2 OFFSET $3;

-- name: ListActiveLoans :many
SELECT * FROM loans
WHERE status = 'active'
ORDER BY due_date
LIMIT $1 OFFSET $2;

-- name: ListOverdueLoans :many
SELECT * FROM loans
WHERE status = 'overdue'
ORDER BY due_date
LIMIT $1 OFFSET $2;

-- name: CreateLoan :one
INSERT INTO loans (
  user_id, book_id, borrowed_date, due_date, status
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: UpdateLoan :one
UPDATE loans
SET 
  user_id = $2,
  book_id = $3,
  borrowed_date = $4,
  due_date = $5,
  returned_date = $6,
  status = $7,
  updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;

-- name: UpdateLoanStatus :one
UPDATE loans
SET 
  status = $2,
  returned_date = $3,
  updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;

-- name: DeleteLoan :exec
DELETE FROM loans
WHERE id = $1;
