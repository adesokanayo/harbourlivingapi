-- name: CreateVenue :one
INSERT INTO venues (
    name,
    address,
    postal_code,
    city,
    province,
    country_code,
    venue_owner,
    banner_image,
    longitude, 
    latitude,
    rating,
    status
) VALUES
    ($1, $2, $3, $4, $5, $6,$7, $8,$9,$10,$11,$12) RETURNING *;

-- name: GetVenue :one
SELECT * FROM venues
WHERE id = $1 LIMIT 1;

-- name: GetAllVenues :many
SELECT * FROM venues
ORDER  by id desc;

-- name: DeleteVenue :exec
DELETE FROM venues
WHERE id = $1;

-- name: RateVenue :exec
UPDATE venues SET rating = $1
where id = $2;

-- name: CreateVenueFavorite :one
INSERT INTO venues_favorites (
    venue_id,
    user_id
) VALUES
  (@venue_id, @user_id) RETURNING *;

-- name: GetFavoriteVenues :many
SELECT * FROM venues_favorites 
where user_id = @user_id
ORDER BY id desc;
  
-- name: DeleteFavoriteVenue :exec
DELETE FROM venues_favorites 
WHERE venue_id = @venue_id 
AND user_id = @user_id;

-- name: UpdateVenue :one
UPDATE venues SET
    name = CASE WHEN @name_to_update::boolean
        THEN @name::text ELSE name END, 
    address = CASE WHEN @address_to_update::boolean
        THEN @address::text ELSE address END,
    postal_code = CASE WHEN @postal_code_to_update::boolean
        THEN @postal_code::text ELSE postal_code END,
    city = CASE WHEN @city_to_update::boolean
        THEN @city::text ELSE city END,
    province = CASE WHEN @province_to_update::boolean
        THEN @province::text ELSE province END,
    country_code = CASE WHEN @country_to_update::boolean
        THEN @country_code::text ELSE country_code END,
    banner_image = CASE WHEN @banner_image_to_update::boolean
        THEN @banner_image::text ELSE banner_image END,
    longitude = CASE WHEN @longitude_to_update::boolean
        THEN @longitude::float ELSE longitude END,
    latitude = CASE WHEN @latitude_to_update::boolean
        THEN @latitude::float ELSE latitude END,
    rating = CASE WHEN @rating_to_update::boolean
        THEN @rating::int ELSE rating END,
    status = CASE WHEN @status_to_update::boolean
        THEN @status::INTEGER ELSE status END
    WHERE id= @id RETURNING *;