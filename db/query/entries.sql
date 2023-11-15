-- name: CreateEntries :one
INSERT INTO entries (
    account_id,
    amount
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetEntries :one
SELECT * FROM entries
WHERE id = $1 LIMIT 1;

-- name: ListEntries :one
SELECT * FROM entries
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateEntry :one
UPDATE entries
SET amount = $2
WHERE id = $1
RETURNING *;

-- name DeleteEntry :exec

DELETE  FROM entries
WHERE id = $1;