-- name: CreateSponsor :one
INSERT INTO sponsor (
user_id
) VALUES
    ($1) RETURNING *;

-- name: GetSponsor :one
SELECT * from sponsor
WHERE id = $1;

-- name: DeleteSponsor :exec
DELETE from sponsor
WHERE id = $1;

-- name: UpdateSponsor :exec
UPDATE events_sponsor
set event_id= $1
WHERE id=$2;

-- name: GetSponsorByEvent :many
SELECT * from events_sponsor
WHERE event_id = $1;

-- name: LinkSponsorToEvent :one
INSERT INTO events_sponsor(
sponsor_id,event_id) 
VALUES($1, $2 ) RETURNING *;
