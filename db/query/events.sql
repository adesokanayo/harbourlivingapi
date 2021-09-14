-- name: GetEvent :one
SELECT * FROM events
WHERE id = $1 LIMIT 1;

-- name: GetEvents :many
SELECT * FROM events e
inner join venue v on  e.venue = v.id
WHERE category = $1
and subcategory =$2
and e.status=$3
ORDER BY e.id desc
LIMIT $4
OFFSET $5;


-- name: DeleteEvent :exec
DELETE FROM events
WHERE id = $1;

-- name: CreateEvent :one
INSERT INTO events (
    title,
    description,
    banner_image,
    start_date,
    end_date,
    venue,
    type,
    user_id,
    category,
    subcategory,
    status,
    image1,
    image2,
    image3,
    video1,
    video2
) VALUES
    ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15,$16) RETURNING *;

