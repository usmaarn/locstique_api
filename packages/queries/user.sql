-- name: FindUserByID :one
SELECT * FROM users WHERE id = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (id, name, email, phone, avatar, type, password, status)
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
    RETURNING *;

-- name: FindUserByEmailAddress :one
SELECT * FROM users WHERE email = $1 LIMIT 1;