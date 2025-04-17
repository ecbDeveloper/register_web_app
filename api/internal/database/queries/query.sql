-- name: CreateUser :exec
INSERT INTO users (
    id, name, email, cpf, phone_number, age, password
) VALUES (
    gen_random_uuid(), $1, $2, $3, $4, $5, $6
);

-- name: SelectUserLoginCredentials :one
SELECT id, email, password 
FROM users WHERE email = $1;

-- name: GetAllUsers :many
SELECT id, name, email, age, phone_number, cpf
FROM users;

-- name: SelectUser :one
SELECT id, name, email, age, phone_number, cpf
FROM users WHERE id = $1;

-- name: UpdateUser :one
UPDATE users SET name = $1, 
    email = $2,
	cpf = $3,
	age = $4,
	phone_number = $5,
    updated_at = $6
WHERE id = $7
RETURNING *;