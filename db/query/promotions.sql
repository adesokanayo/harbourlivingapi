-- name: GetPromotion :one
SELECT * FROM promotions
WHERE id = @id LIMIT 1;

-- name: GetPromotionForEvent :one
SELECT * FROM promotions
WHERE event_id = @event_id LIMIT 1;

-- name: DeletePromotion :exec
DELETE FROM promotions
WHERE id = $1;

-- name: CreatePromotion :one
INSERT INTO promotions (
    event_id,
    user_id,
    plan_id,
    start_date,
    end_date
) VALUES
  ($1, $2, $3, $4, $5 ) RETURNING *;

-- name: UpdatePromotion :one
UPDATE promotions SET
 event_id = CASE WHEN @event_id_to_update::boolean
        THEN @event_id::int ELSE event_id END, 
 plan_id = CASE WHEN @plan_id_to_update::boolean
        THEN @plan_id::int ELSE plan_id END,
 start_date = CASE WHEN @start_date_to_update::boolean
        THEN @start_date::timestamptz ELSE start_date END,
 end_date = CASE WHEN @end_date_to_update::boolean
        THEN @end_date::timestamptz ELSE end_date END
WHERE id = @id RETURNING *;