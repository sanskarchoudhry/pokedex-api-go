-- name: CreateType :one
INSERT INTO types (id, name)
VALUES ($1, $2)
ON CONFLICT (id) DO UPDATE
SET name = EXCLUDED.name
RETURNING *;

-- name: ListTypes :many
SELECT * FROM types
ORDER BY id;