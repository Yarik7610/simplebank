-- name: CreteEntry :one
INSERT INTO entries (account_id, amount)
VALUES ($1, $2)
RETURNING *;

-- name: GetEntry :one
SELECT * FROM entries
WHERE id = $1;

-- name: ListEntries :many
SELECT * FROM entries
ORDER BY id
FETCH FIRST $1 ROWS ONLY
OFFSET $2 ROWS;