-- name: CreateUser :one
INSERT INTO users (
    id, name, email, password
) VALUES (
    gen_random_uuid(), $1, $2, $3
) 
    RETURNING id;

-- name: SelectUserLoginCredentials :one
SELECT id, email, password 
FROM users WHERE email = $1;

-- name: GetAllUsers :many
SELECT id, name, email 
FROM users;

-- name: SelectUser :one
SELECT id, name, email 
FROM users WHERE id = $1;

-- name: UpdateUser :one
UPDATE users SET name = $1, 
    email = $2,
    password = $3,
    updated_at = $4 
WHERE id = $5
RETURNING *;