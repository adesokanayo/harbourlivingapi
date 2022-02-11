-- name: GetEvent :one
SELECT * FROM events
WHERE id = $1 LIMIT 1;


-- name: GetEvents :many
SELECT * FROM events
WHERE status = sqlc.arg('status') AND
 end_date >= CURRENT_DATE
  AND (CASE WHEN sqlc.arg('categoryFilter')::bool THEN category = sqlc.arg('category') ELSE TRUE END)
  AND (CASE WHEN sqlc.arg('titleFilter')::bool THEN title ILIKE sqlc.arg('title') ELSE TRUE END)
ORDER BY 
  CASE WHEN sqlc.arg('startDateAsc')::bool THEN start_date END asc,
  CASE WHEN sqlc.arg('startDateDesc')::bool THEN start_date END desc,
  CASE WHEN sqlc.arg('endDateAsc')::bool THEN end_date END asc,
  CASE WHEN sqlc.arg('endDateDesc')::bool THEN end_date END desc,
  CASE WHEN sqlc.arg('defaultOrder')::bool THEN id END desc
LIMIT sqlc.arg('limit')
OFFSET sqlc.arg('offset') ROWS;

-- name: GetEventsFilter :many
SELECT * FROM events e
WHERE e.status = $1 AND
e.category = $2 AND
e.end_date >= CURRENT_DATE
ORDER BY e.id desc
LIMIT $3
OFFSET $4 ROWS;

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
    status
) VALUES
  ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING *;

-- name: UpdateEvent :one
UPDATE events SET
    title = CASE WHEN @title_to_update::boolean
        THEN @title::text ELSE title END, 
    description = CASE WHEN @description_to_update::boolean
        THEN @description::text ELSE description END,
    banner_image = CASE WHEN @banner_image_to_update::boolean
        THEN @banner_image::text ELSE banner_image END,
    start_date = CASE WHEN @start_date_to_update::boolean
        THEN @start_date::timestamptz ELSE start_date END,
    end_date =CASE WHEN @end_date_to_update::boolean
        THEN @end_date::timestamptz ELSE end_date END,
    venue = CASE WHEN @venue_to_update::boolean
        THEN @venue::INTEGER ELSE venue END,
    type = CASE WHEN @type_to_update::boolean
        THEN @type::INTEGER ELSE type END,
    user_id = CASE WHEN @user_id_to_update::boolean
        THEN @user_id::INTEGER ELSE user_id END,
    category = CASE WHEN @category_to_update::boolean
        THEN @category::INTEGER ELSE category END,
    status = CASE WHEN @status_to_update::boolean
        THEN @status::INTEGER ELSE status END
    WHERE id= @id RETURNING *;

-- name: UpdateEventStatus :one
UPDATE events
set status = $1
where id = $2 RETURNING events.Id, events.status;

-- name: GetEventsByLocation :many 
SELECT *, point($1,$2) <@>  (point(v.longitude, v.latitude)::point) as distance
FROM venues v, events e
WHERE (point($1,$2) <@> point(longitude, latitude)) < $3  
ORDER BY distance desc;

-- name: CreateFavoriteEvent :one
INSERT INTO events_favorites (
    event_id,
    user_id
) VALUES
  (@event_id, @user_id) RETURNING *;

-- name: GetFavoriteEvents :many
SELECT * FROM events_favorites 
where user_id = @user_id
ORDER BY id desc;

-- name: DeleteFavoriteEvent :exec
DELETE FROM events_favorites 
WHERE event_id = @event_id 
AND user_id = @user_id;


-- name: CreateViewEvent :one
INSERT INTO events_views (
    event_id,
    user_id
) VALUES
  (@event_id, @user_id) RETURNING *;

-- name: GetViewedEvents :many
SELECT * FROM events_views 
where user_id = @user_id
ORDER BY id desc;