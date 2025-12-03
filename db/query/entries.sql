-- name: CreateEntry :one
INSERT INTO entries (
    account_id,
    amount
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetEntry :one
SELECT * FROM entries
WHERE id = $1 LIMIT 1;

-- name: GetEntryAnAccount :one
SELECT e.id AS entry_id,
       e.amount,
       a.id AS account_id,
       a.owner,
       a.balance  
FROM entries e JOIN accounts a ON a.id = e.account_id WHERE e.id = $1;

-- name: ListEntriesByAccount :many
SELECT * FROM entries WHERE account_id = $1 ORDER BY id LIMIT $2 OFFSET $3;

-- name: UpdateEntryAmount :one
UPDATE entries SET amount = $2 WHERE id = $1 RETURNING *;

-- name: DeleteEntry :exec
DELETE FROM entries
WHERE id = $1;