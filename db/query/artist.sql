-- name: CreateArtist :one
INSERT INTO artists (
user_id
) VALUES
    ($1) RETURNING *;

-- name: GetArtist :one
SELECT * from artists
WHERE id = $1;

-- name: DeleteArtist :exec
DELETE from artists
WHERE id = $1;

-- name: UpdateEventArtist :exec
UPDATE events_artists
set event_id= $1
WHERE id=$2;

-- name: UpdateArtist :one
UPDATE artists SET
 display_name = CASE WHEN @display_name_to_update::boolean
        THEN @display_name::text ELSE display_name END, 
 avatar_url = CASE WHEN @avatar_url_to_update::boolean
        THEN @avatar_url::text ELSE avatar_url END,
 short_bio = CASE WHEN @short_bio_to_update::boolean
        THEN @short_bio::text ELSE short_bio END
WHERE id = @id RETURNING *;

-- name: GetArtistByEvent :many
SELECT * from events_artists
WHERE event_id = $1;

-- name: LinkArtistToEvent :one
INSERT INTO events_artists(
artist_id,event_id)
VALUES($1, $2 ) RETURNING *;
