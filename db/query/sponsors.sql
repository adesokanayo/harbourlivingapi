-- name: CreateSponsor :one
INSERT INTO sponsors (
user_id
) VALUES
    ($1) RETURNING *;

-- name: GetSponsor :one
SELECT * from sponsors
WHERE id = $1;

-- name: DeleteSponsor :exec
DELETE from sponsors
WHERE id = $1;

-- name: UpdateSponsor :exec
UPDATE events_sponsors
set event_id= $1
WHERE id=$2;

-- name: GetSponsorByEvent :many
SELECT * from events_sponsors
WHERE event_id = $1;

-- name: LinkSponsorToEvent :one
INSERT INTO events_sponsors(
sponsor_id,event_id) 
VALUES($1, $2 ) RETURNING *;
