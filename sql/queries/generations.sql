-- name: CreateGeneration :one
INSERT INTO generations (id, name, region_name)
VALUES ($1, $2, $3)
ON CONFLICT (id) DO UPDATE
SET name = EXCLUDED.name, 
    region_name = EXCLUDED.region_name
RETURNING *;

-- name: ListGenerations :many
SELECT * FROM generations
ORDER BY id;