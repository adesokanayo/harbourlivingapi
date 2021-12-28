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

-- name: UpdateSponsor :one
UPDATE sponsors SET
 display_name = CASE WHEN @display_name_to_update::boolean
        THEN @display_name::text ELSE display_name END, 
 avatar_url = CASE WHEN @avatar_url_to_update::boolean
        THEN @avatar_url::text ELSE avatar_url END,
 short_bio = CASE WHEN @short_bio_to_update::boolean
        THEN @short_bio::text ELSE short_bio END
WHERE id = @id RETURNING *;

-- name: GetSponsorByEvent :many
SELECT * from events_sponsors
WHERE event_id = $1;

-- name: LinkSponsorToEvent :one
INSERT INTO events_sponsors(
sponsor_id,event_id) 
VALUES($1, $2 ) RETURNING *;
