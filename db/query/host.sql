-- name: CreateHost :one
INSERT INTO host (
user_id
) VALUES
    ($1) RETURNING *;

-- name: GetHost :one
SELECT * from host
WHERE id = $1;

-- name: DeleteHost :exec
DELETE from host
WHERE id = $1;

-- name: UpdateHost :exec
UPDATE events_host
set event_id= $1
WHERE id=$2;

-- name: GetHostByEvent :many
SELECT * from events_host
WHERE event_id = $1;

-- name: LinkHostToEvent :one
INSERT INTO events_host(
host_id,event_id)
VALUES($1, $2 ) RETURNING *;
