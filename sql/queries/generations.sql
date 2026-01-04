-- name: CreateGeneration :one
INSERT INTO generations (id, name, region_name)
VALUES ($1, $2, $3)
ON CONFLICT (id) DO NOTHING
RETURNING *;

-- name: ListGenerations :many
SELECT * FROM generations
ORDER BY id;