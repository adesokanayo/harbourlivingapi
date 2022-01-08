// Code generated by sqlc. DO NOT EDIT.
// source: promotions.sql

package db

import (
	"context"
	"time"
)

const createPromotion = `-- name: CreatePromotion :one
INSERT INTO promotions (
    event_id,
    user_id,
    plan_id,
    start_date,
    end_date
) VALUES
  ($1, $2, $3, $4, $5 ) RETURNING id, event_id, user_id, plan_id, start_date, end_date, created_at
`

type CreatePromotionParams struct {
	EventID   int32     `json:"event_id"`
	UserID    int32     `json:"user_id"`
	PlanID    int32     `json:"plan_id"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

func (q *Queries) CreatePromotion(ctx context.Context, arg CreatePromotionParams) (Promotion, error) {
	row := q.db.QueryRowContext(ctx, createPromotion,
		arg.EventID,
		arg.UserID,
		arg.PlanID,
		arg.StartDate,
		arg.EndDate,
	)
	var i Promotion
	err := row.Scan(
		&i.ID,
		&i.EventID,
		&i.UserID,
		&i.PlanID,
		&i.StartDate,
		&i.EndDate,
		&i.CreatedAt,
	)
	return i, err
}

const deletePromotion = `-- name: DeletePromotion :exec
DELETE FROM promotions
WHERE id = $1
`

func (q *Queries) DeletePromotion(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deletePromotion, id)
	return err
}

const getAllPromotions = `-- name: GetAllPromotions :many
SELECT id, event_id, user_id, plan_id, start_date, end_date, created_at  FROM promotions
order by id desc
`

func (q *Queries) GetAllPromotions(ctx context.Context) ([]Promotion, error) {
	rows, err := q.db.QueryContext(ctx, getAllPromotions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Promotion{}
	for rows.Next() {
		var i Promotion
		if err := rows.Scan(
			&i.ID,
			&i.EventID,
			&i.UserID,
			&i.PlanID,
			&i.StartDate,
			&i.EndDate,
			&i.CreatedAt,
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

const getPromotion = `-- name: GetPromotion :one
SELECT id, event_id, user_id, plan_id, start_date, end_date, created_at FROM promotions
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetPromotion(ctx context.Context, id int32) (Promotion, error) {
	row := q.db.QueryRowContext(ctx, getPromotion, id)
	var i Promotion
	err := row.Scan(
		&i.ID,
		&i.EventID,
		&i.UserID,
		&i.PlanID,
		&i.StartDate,
		&i.EndDate,
		&i.CreatedAt,
	)
	return i, err
}

const getPromotionForEvent = `-- name: GetPromotionForEvent :one
SELECT id, event_id, user_id, plan_id, start_date, end_date, created_at FROM promotions
WHERE event_id = $1 LIMIT 1
`

func (q *Queries) GetPromotionForEvent(ctx context.Context, eventID int32) (Promotion, error) {
	row := q.db.QueryRowContext(ctx, getPromotionForEvent, eventID)
	var i Promotion
	err := row.Scan(
		&i.ID,
		&i.EventID,
		&i.UserID,
		&i.PlanID,
		&i.StartDate,
		&i.EndDate,
		&i.CreatedAt,
	)
	return i, err
}

const updatePromotion = `-- name: UpdatePromotion :one
UPDATE promotions SET
 event_id = CASE WHEN $1::boolean
        THEN $2::int ELSE event_id END, 
 plan_id = CASE WHEN $3::boolean
        THEN $4::int ELSE plan_id END,
 start_date = CASE WHEN $5::boolean
        THEN $6::timestamptz ELSE start_date END,
 end_date = CASE WHEN $7::boolean
        THEN $8::timestamptz ELSE end_date END
WHERE id = $9 RETURNING id, event_id, user_id, plan_id, start_date, end_date, created_at
`

type UpdatePromotionParams struct {
	EventIDToUpdate   bool      `json:"event_id_to_update"`
	EventID           int32     `json:"event_id"`
	PlanIDToUpdate    bool      `json:"plan_id_to_update"`
	PlanID            int32     `json:"plan_id"`
	StartDateToUpdate bool      `json:"start_date_to_update"`
	StartDate         time.Time `json:"start_date"`
	EndDateToUpdate   bool      `json:"end_date_to_update"`
	EndDate           time.Time `json:"end_date"`
	ID                int32     `json:"id"`
}

func (q *Queries) UpdatePromotion(ctx context.Context, arg UpdatePromotionParams) (Promotion, error) {
	row := q.db.QueryRowContext(ctx, updatePromotion,
		arg.EventIDToUpdate,
		arg.EventID,
		arg.PlanIDToUpdate,
		arg.PlanID,
		arg.StartDateToUpdate,
		arg.StartDate,
		arg.EndDateToUpdate,
		arg.EndDate,
		arg.ID,
	)
	var i Promotion
	err := row.Scan(
		&i.ID,
		&i.EventID,
		&i.UserID,
		&i.PlanID,
		&i.StartDate,
		&i.EndDate,
		&i.CreatedAt,
	)
	return i, err
}
