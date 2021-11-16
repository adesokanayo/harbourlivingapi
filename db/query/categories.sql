-- name: GetCategory :one
SELECT * FROM categories
WHERE id = $1 LIMIT 1;

-- name: GetCategories :many
SELECT * FROM categories;

-- name: UpdateCategoryStatus :exec
UPDATE categories
set status = $1
where id = $2;

--name: DeleteCategory :exec
DELETE  from categories
where id =$1 ;

-- name: GetSubCategory :one
SELECT * FROM subcategories
WHERE category_id = $1 LIMIT 1;


-- name: GetSubCategories :many
SELECT * FROM subcategories
WHERE category_id = $1;

-- name: UpdateSubCategoryStatus :exec
UPDATE categories
set status = $1
where id = $2;

--name: DeleteSubCategory :exec
DELETE  from Subcategories
where id =$1;