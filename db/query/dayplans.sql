-- name: GetDayplan :one
SELECT * FROM dayplans
WHERE id = @id LIMIT 1;

-- name: GetDayplanForSchedules :many
SELECT * FROM dayplans 
WHERE schedule_id = @schedule_id; 

-- name: DeleteDayPlan :exec
DELETE FROM dayplans
WHERE id = $1;

-- name: GetAllDayplans :many
SELECT *  FROM dayplans
order by id desc;

-- name: CreateDayplan :one
INSERT INTO dayplans (
    start_time,
    end_time,
    schedule_id,
    title,
    description,
    performer_name
) VALUES
  ($1, $2, $3, $4, $5, $6 ) RETURNING *;

-- name: UpdateDayPlan :one
UPDATE dayplans SET
 start_time = CASE WHEN @start_time_to_update::boolean
        THEN @start_time::timestamp ELSE start_time END, 
 end_time = CASE WHEN @end_time_to_update::boolean
        THEN @end_time::timestamp ELSE end_time END,
 title = CASE WHEN @title_to_update::boolean
        THEN @title::text ELSE title END,
 description = CASE WHEN @description_to_update::boolean
        THEN @description::text ELSE description END,
 performer_name = CASE WHEN @performer_name_to_update::boolean
        THEN @performer_name::text ELSE performer_name END
WHERE id = @id RETURNING *;

