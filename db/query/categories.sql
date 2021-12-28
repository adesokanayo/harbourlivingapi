-- name: GetCategory :one
SELECT * FROM categories
WHERE id = $1 LIMIT 1;

-- name: GetCategories :many
SELECT * FROM categories;

-- name: UpdateCategoryStatus :exec
UPDATE categories
set status = $1
where id = $2;

-- name: DeleteCategory :exec
DELETE  from categories
where id =$1 ;

-- name: CreateCategory :one
INSERT INTO categories (
    description,
    image,
    status
) VALUES
  (@description, @image, @status) RETURNING *;