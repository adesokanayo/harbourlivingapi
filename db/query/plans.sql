-- name: CreatePlan :one
INSERT INTO plans (
name,
description,
price
) VALUES
(@name, @description, @price) RETURNING *;

-- name: GetPlan :one
SELECT * from plans
WHERE id = $1;

-- name: DeletePlan :exec
DELETE from plans
WHERE id = $1;
  
-- name: UpdatePlan :one
UPDATE plans SET
 name = CASE WHEN @name_to_update::boolean
        THEN @name::text ELSE name END, 
 description = CASE WHEN @description_to_update::boolean
        THEN @description::text ELSE description END,
 price = CASE WHEN @price_to_update::boolean
        THEN @price::float ELSE price END
WHERE id = @id RETURNING *;

