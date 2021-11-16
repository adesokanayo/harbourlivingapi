-- name: CreateHost :one
INSERT INTO hosts (
user_id
) VALUES
    ($1) RETURNING *;

-- name: GetHost :one
SELECT * from hosts
WHERE id = $1;

-- name: DeleteHost :exec
DELETE from hosts
WHERE id = $1;

-- name: UpdateHost :exec
UPDATE events_hosts
set event_id= $1
WHERE id=$2;

-- name: GetHostByEvent :many
SELECT * from events_hosts
WHERE event_id = $1;

-- name: LinkHostToEvent :one
INSERT INTO events_hosts(
host_id,event_id)
VALUES($1, $2 ) RETURNING *;
