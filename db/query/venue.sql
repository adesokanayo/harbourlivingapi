-- name: CreateVenue :one
INSERT INTO venue (
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
SELECT * FROM venue
WHERE id = $1 LIMIT 1;

-- name: GetAllVenues :many
SELECT * FROM venue
ORDER  by id;

-- name: DeleteVenue :exec
DELETE FROM venue
WHERE id = $1;

-- name: CreateVirtualVenue :one
INSERT INTO venue (
    name,
    url,
    virtual
) VALUES
    ($1, $2, $3) RETURNING *;

-- name: RateVenue :exec
UPDATE venue SET rating = $1
where id = $2;