-- name: CreateUser :one
INSERT INTO users (
    phone,
    first_name,
    last_name,
    email,
    password,
    username,
    usertype,
    date_of_birth,
    avatar_url,
    activation_code
) VALUES
    ($1, $2, $3, $4, $5, $6, $7,$8,$9, $10) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUsername :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: GetEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: GetAllUsers :many
SELECT * FROM users
ORDER  by id;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: UpdateUser :one
UPDATE users SET
    phone = CASE WHEN @phone_to_update::boolean
        THEN @phone::text ELSE phone END, 
    first_name = CASE WHEN @first_name_to_update::boolean
        THEN @first_name::text ELSE first_name END,
    last_name = CASE WHEN @last_name_to_update::boolean
        THEN @last_name::text ELSE last_name END,
    email = CASE WHEN @email_to_update::boolean
        THEN @email::text ELSE email END,
    password =CASE WHEN @password_to_update::boolean
        THEN @password::text ELSE password END,
    username = CASE WHEN @username_to_update::boolean
        THEN @username::text ELSE username END,
    date_of_birth = CASE WHEN @date_of_birth_to_update::boolean
        THEN @date_of_birth::timestamp ELSE date_of_birth END,
    avatar_url = CASE WHEN @avatar_url_to_update::boolean
        THEN @avatar_url::text ELSE avatar_url END
    WHERE id= @id RETURNING *;