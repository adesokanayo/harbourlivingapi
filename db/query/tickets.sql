-- name: CreateTicket :one
INSERT INTO tickets (
  name,
  event_id,
  quantity, 
  price, 
  status,
  currency
) VALUES
    ($1, $2, $3, $4, $5, $6) RETURNING *;

-- name: GetTicket :one
SELECT * FROM tickets
WHERE id = $1 LIMIT 1;

-- name: GetTicketsByEventID :many
SELECT * FROM tickets
where event_id = $1;

-- name: GetAllTickets :many
SELECT * FROM tickets
ORDER  by id;

-- name: DeleteTicket :exec
DELETE FROM tickets
WHERE id = $1;
