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
    avatar_url
) VALUES
    ($1, $2, $3, $4, $5, $6, $7,$8,$9) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUsername :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;
-- name: GetAllUsers :many
SELECT * FROM users
ORDER  by id;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;