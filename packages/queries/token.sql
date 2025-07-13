-- name: CreateToken :one
INSERT INTO tokens (id, user_id, expires_at, user_agent, ip_address)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING *;

-- name: FindTokenByID :one
SELECT * FROM tokens WHERE id = $1;

-- name: DeleteToken :exec
DELETE FROM tokens WHERE id = $1;

-- name: UpdateToken :exec
UPDATE tokens SET expires_at = $1 WHERE id = $2;


