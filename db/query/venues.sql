-- name: CreateVenue :one
INSERT INTO venues (
    name,
    address,
    postal_code,
    city,
    province,
    country_code,
    url,
    virtual,
    longitude, 
    latitude
) VALUES
    ($1, $2, $3, $4, $5, $6,$7, $8,$9,$10) RETURNING *;

-- name: GetVenue :one
SELECT * FROM venues
WHERE id = $1 LIMIT 1;

-- name: GetAllVenues :many
SELECT * FROM venues
ORDER  by id;

-- name: DeleteVenue :exec
DELETE FROM venues
WHERE id = $1;

-- name: CreateVirtualVenue :one
INSERT INTO venues (
    name,
    url,
    virtual
) VALUES
    ($1, $2, $3) RETURNING *;

-- name: RateVenue :exec
UPDATE venues SET rating = $1
where id = $2;