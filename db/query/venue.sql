-- name: CreateVenue :one
INSERT INTO venue (
    name,
    address,
    postal_code,
    city,
    province,
    country_code
) VALUES
    ($1, $2, $3, $4, $5, $6) RETURNING *;

-- name: GetVenue :one
SELECT * FROM venue
WHERE id = $1 LIMIT 1;

-- name: GetAllVenues :many
SELECT * FROM venue
ORDER  by id;

-- name: DeleteVenue :exec
DELETE FROM venue
WHERE id = $1;
