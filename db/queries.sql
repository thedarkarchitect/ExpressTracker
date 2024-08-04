-- name: CreateUser :one
INSERT INTO users (email, username, role, password) VALUES ($1, $2, $3, $4) RETURNING id, email, username, role, password, created_at;

-- name: GetUserByUsername :one
SELECT id, email, username, role, password, created_at FROM users WHERE username = $1;

-- name: CreateExpense :one
INSERT INTO expenses (user_id, description, amount) VALUES ($1, $2, $3) RETURNING id, user_id, description, amount, date;

-- name: GetExpensesByUserID :many
SELECT id, user_id, description, amount, date FROM expenses WHERE user_id = $1;

-- name: UpdateExpense :exec
UPDATE expenses SET description = $1, amount = $2 WHERE id = $3 AND user_id = $4;

-- name: DeleteExpense :exec
DELETE FROM expenses WHERE id = $1;