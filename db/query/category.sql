-- name: GetCategory :one
SELECT * FROM category
WHERE id = $1 LIMIT 1;

-- name: GetCategories :many
SELECT * FROM category;

-- name: UpdateCategoryStatus :exec
UPDATE Category
set status = $1
where id = $2;

--name: DeleteCategory :exec
DELETE  from Category
where id =$1 ;

-- name: GetSubCategory :one
SELECT * FROM subcategory
WHERE category_id = $1 LIMIT 1;


-- name: GetSubCategories :many
SELECT * FROM subcategory
WHERE category_id = $1;

-- name: UpdateSubCategoryStatus :exec
UPDATE Category
set status = $1
where id = $2;

--name: DeleteSubCategory :exec
DELETE  from Subcategory
where id =$1;