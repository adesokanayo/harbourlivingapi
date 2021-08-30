-- name: CreateTicket :one
INSERT INTO ticket (
  name,
  event_id,
  quantity, 
  price, 
  status
) VALUES
    ($1, $2, $3, $4,$5) RETURNING *;

-- name: GetTicket :one
SELECT * FROM ticket
WHERE id = $1 LIMIT 1;

-- name: GetAllTickets :many
SELECT * FROM ticket
ORDER  by id;

-- name: DeleteTicket :exec
DELETE FROM ticket
WHERE id = $1;
