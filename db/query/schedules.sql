-- name: GetSchedule :one
SELECT * FROM schedules
WHERE id = @id LIMIT 1;

-- name: GetSchedulesForEvent :many
SELECT * FROM schedules
WHERE event_id = @event_id; 

-- name: DeleteSchedule :exec
DELETE FROM schedules
WHERE id = $1;

-- name: GetAllSchedule :many
SELECT *  FROM schedules
order by id desc;

-- name: CreateSchedule :one
INSERT INTO schedules (
    event_id,
    date,
    start_time,
    end_time
) VALUES
  ($1, $2, $3, $4 ) RETURNING *;

-- name: UpdateSchedule :one
UPDATE schedules SET
 event_id = CASE WHEN @event_id_to_update::boolean
        THEN @event_id::int ELSE event_id END, 
 date = CASE WHEN @date_id_to_update::boolean
        THEN @date::timestamp ELSE date END,
 start_time = CASE WHEN @start_time_to_update::boolean
        THEN @start_time::timestamp ELSE start_time END,
 end_time = CASE WHEN @end_time_to_update::boolean
        THEN @end_time::timestamp ELSE end_time END
WHERE id = @id RETURNING *;