-- name: CreateUser :one
INSERT INTO users (id, username, email, password, roles, xp, is_banned, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING id, username, email, password, roles, xp, is_banned, created_at, updated_at;

-- name: GetAllUsers :many
SELECT id, username, email, password, roles, xp, is_banned, created_at, updated_at FROM users;

-- name: GetUserById :one
SELECT id, username, email, password, roles, xp, is_banned, created_at, updated_at FROM users WHERE id = $1;

-- name: GetUserByEmail :one
SELECT id, username, email, password, roles, xp, is_banned, created_at, updated_at FROM users WHERE email = $1;

-- name: UpdateUserById :one
UPDATE users
SET username = COALESCE($2, username), 
    email = COALESCE($3, email), 
    password = COALESCE($4, password), 
    roles = COALESCE($5, roles), 
    xp = COALESCE($6, xp), 
    is_banned = COALESCE($7, is_banned), 
    updated_at = NOW()
WHERE id = $1
RETURNING id, username, email, password, roles, xp, is_banned, created_at, updated_at;

-- name: DeleteUserById :exec
DELETE FROM users WHERE id = $1;
