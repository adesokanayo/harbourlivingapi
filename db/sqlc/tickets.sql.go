// Code generated by sqlc. DO NOT EDIT.
// source: tickets.sql

package db

import (
	"context"
	"database/sql"
)

const createTicket = `-- name: CreateTicket :one
INSERT INTO tickets (
  name,
  event_id,
  price, 
  currency,
  description
) VALUES
    ($1, $2, $3, $4, $5) RETURNING id, name, event_id, price, currency, description
`

type CreateTicketParams struct {
	Name        string         `json:"name"`
	EventID     int32          `json:"event_id"`
	Price       float64        `json:"price"`
	Currency    string         `json:"currency"`
	Description sql.NullString `json:"description"`
}

func (q *Queries) CreateTicket(ctx context.Context, arg CreateTicketParams) (Ticket, error) {
	row := q.db.QueryRowContext(ctx, createTicket,
		arg.Name,
		arg.EventID,
		arg.Price,
		arg.Currency,
		arg.Description,
	)
	var i Ticket
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.EventID,
		&i.Price,
		&i.Currency,
		&i.Description,
	)
	return i, err
}

const deleteTicket = `-- name: DeleteTicket :exec
DELETE FROM tickets
WHERE id = $1
`

func (q *Queries) DeleteTicket(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteTicket, id)
	return err
}

const getAllTickets = `-- name: GetAllTickets :many
SELECT id, name, event_id, price, currency, description FROM tickets
ORDER  by id
`

func (q *Queries) GetAllTickets(ctx context.Context) ([]Ticket, error) {
	rows, err := q.db.QueryContext(ctx, getAllTickets)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Ticket{}
	for rows.Next() {
		var i Ticket
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.EventID,
			&i.Price,
			&i.Currency,
			&i.Description,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTicket = `-- name: GetTicket :one
SELECT id, name, event_id, price, currency, description FROM tickets
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetTicket(ctx context.Context, id int32) (Ticket, error) {
	row := q.db.QueryRowContext(ctx, getTicket, id)
	var i Ticket
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.EventID,
		&i.Price,
		&i.Currency,
		&i.Description,
	)
	return i, err
}

const getTicketsByEventID = `-- name: GetTicketsByEventID :many
SELECT id, name, event_id, price, currency, description FROM tickets
where event_id = $1
`

func (q *Queries) GetTicketsByEventID(ctx context.Context, eventID int32) ([]Ticket, error) {
	rows, err := q.db.QueryContext(ctx, getTicketsByEventID, eventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Ticket{}
	for rows.Next() {
		var i Ticket
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.EventID,
			&i.Price,
			&i.Currency,
			&i.Description,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
