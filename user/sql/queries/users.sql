-- name: CreateUser :one
INSERT INTO users (id, name, email, password, level, badges, is_banned, created_at, updated_at)
VALUES (uuid_generate_v4(), $1, $2, $3, $4, $5, $6, NOW(), NOW())
RETURNING *;

-- name: GetAllUsers :many
SELECT * FROM users;

-- name: GetUserById :one
SELECT * FROM users WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: UpdateUser :one
UPDATE users
SET name = $2, email = $3, password = $4, level = $5, badges = $6, is_banned = $7, updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;
