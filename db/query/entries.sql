-- name: CreateEntry :one
INSERT INTO entries (account_id, amount)
VALUES ($1, $2)
returning *;

-- name: GetEntry :one
SELECT *
from entries
where id = $1
limit 1;

-- name: ListEntries :many
select *
from entries
order by id
limit $1 offset $2;

-- name: UpdateEntry :one
update entries
set amount = $2
where id = $1
returning *;

-- name: DeleteEntry :exec
delete
from entries
where id = $1;