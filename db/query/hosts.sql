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

-- name: UpdateEventHost :exec
UPDATE events_hosts
set event_id= $1
WHERE id=$2;

-- name: UpdateHost :one
UPDATE hosts SET
 display_name = CASE WHEN @display_name_to_update::boolean
        THEN @display_name::text ELSE display_name END, 
 avatar_url = CASE WHEN @avatar_url_to_update::boolean
        THEN @avatar_url::text ELSE avatar_url END,
 short_bio = CASE WHEN @short_bio_to_update::boolean
        THEN @short_bio::text ELSE short_bio END
WHERE id = @id RETURNING *;

-- name: GetHostByEvent :many
SELECT * from events_hosts
WHERE event_id = $1;

-- name: LinkHostToEvent :one
INSERT INTO events_hosts(
host_id,event_id)
VALUES($1, $2 ) RETURNING *;
