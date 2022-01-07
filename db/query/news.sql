-- name: GetNews :one
SELECT * FROM news
WHERE id = @id LIMIT 1;

-- name: DeleteNews :exec
DELETE FROM news
WHERE id = $1;

-- name: CreateNews :one
INSERT INTO news (
    title,
    description,
    feature_image,
    body,
    user_id,
    publish_date,
    tags
) VALUES
  ($1, $2, $3, $4, $5, $6, $7 ) RETURNING *;

-- name: UpdateNews :one
UPDATE news SET
 title = CASE WHEN @title_id_to_update::boolean
        THEN @title::text ELSE title END, 
 description = CASE WHEN @description_to_update::boolean
        THEN @description::text ELSE description END,
 feature_image = CASE WHEN @feature_image_to_update::boolean
        THEN @feature_image::text ELSE feature_image END,
 body = CASE WHEN @body_to_update::boolean
        THEN @body::text ELSE body END,
 publish_date = CASE WHEN @publish_date_to_update::boolean
        THEN @publish_date::timestamptz ELSE publish_date END,
 tags = CASE WHEN @tags_date_to_update::boolean
        THEN @tags::text ELSE tags END
WHERE id = @id RETURNING *;