-- name: CreatePlan :one
INSERT INTO plans (
name,
description,
price,
no_of_days
) VALUES
(@name, @description, @price, @no_of_days) RETURNING *;

-- name: GetPlan :one
SELECT * from plans
WHERE id = $1;

-- name: GetAllPlans :many
SELECT * from plans
order by id desc;

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
        THEN @price::float ELSE price END,
 no_of_days = CASE WHEN @no_of_days_to_update::boolean
        THEN @no_of_days::int ELSE no_of_days END
WHERE id = @id RETURNING *;

