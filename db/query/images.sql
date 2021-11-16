-- name: GetImage :one
SELECT * FROM images
WHERE id = $1 LIMIT 1;

-- name: GetAllImages :many
SELECT * FROM images;

--name: DeleteImage :exec
DELETE  from images
where id =$1 ;

-- name: GetImagesByEvent :many
SELECT * FROM images a, events_images b
WHERE event_id = $1 and a.id = b.image_id ;

-- name: LinkImageToEvent :exec
INSERT into events_images (event_id, image_id)
VALUES($1,$2); 

-- name: CreateImage :one
INSERT INTO images ( name, url  )
VALUES ($1,$2) RETURNING *;
