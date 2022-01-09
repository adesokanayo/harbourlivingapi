// Code generated by sqlc. DO NOT EDIT.
// source: dayplans.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createDayplan = `-- name: CreateDayplan :one
INSERT INTO dayplans (
    start_time,
    end_time,
    schedule_id,
    title,
    description,
    performer_name
) VALUES
  ($1, $2, $3, $4, $5, $6 ) RETURNING id, start_time, end_time, schedule_id, title, description, performer_name, created_at
`

type CreateDayplanParams struct {
	StartTime     time.Time      `json:"start_time"`
	EndTime       time.Time      `json:"end_time"`
	ScheduleID    int32          `json:"schedule_id"`
	Title         sql.NullString `json:"title"`
	Description   sql.NullString `json:"description"`
	PerformerName sql.NullString `json:"performer_name"`
}

func (q *Queries) CreateDayplan(ctx context.Context, arg CreateDayplanParams) (Dayplan, error) {
	row := q.db.QueryRowContext(ctx, createDayplan,
		arg.StartTime,
		arg.EndTime,
		arg.ScheduleID,
		arg.Title,
		arg.Description,
		arg.PerformerName,
	)
	var i Dayplan
	err := row.Scan(
		&i.ID,
		&i.StartTime,
		&i.EndTime,
		&i.ScheduleID,
		&i.Title,
		&i.Description,
		&i.PerformerName,
		&i.CreatedAt,
	)
	return i, err
}

const deleteDayPlan = `-- name: DeleteDayPlan :exec
DELETE FROM dayplans
WHERE id = $1
`

func (q *Queries) DeleteDayPlan(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteDayPlan, id)
	return err
}

const getAllDayplans = `-- name: GetAllDayplans :many
SELECT id, start_time, end_time, schedule_id, title, description, performer_name, created_at  FROM dayplans
order by id desc
`

func (q *Queries) GetAllDayplans(ctx context.Context) ([]Dayplan, error) {
	rows, err := q.db.QueryContext(ctx, getAllDayplans)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Dayplan{}
	for rows.Next() {
		var i Dayplan
		if err := rows.Scan(
			&i.ID,
			&i.StartTime,
			&i.EndTime,
			&i.ScheduleID,
			&i.Title,
			&i.Description,
			&i.PerformerName,
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

const getDayplan = `-- name: GetDayplan :one
SELECT id, start_time, end_time, schedule_id, title, description, performer_name, created_at FROM dayplans
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetDayplan(ctx context.Context, id int32) (Dayplan, error) {
	row := q.db.QueryRowContext(ctx, getDayplan, id)
	var i Dayplan
	err := row.Scan(
		&i.ID,
		&i.StartTime,
		&i.EndTime,
		&i.ScheduleID,
		&i.Title,
		&i.Description,
		&i.PerformerName,
		&i.CreatedAt,
	)
	return i, err
}

const getDayplanForSchedules = `-- name: GetDayplanForSchedules :many
SELECT id, start_time, end_time, schedule_id, title, description, performer_name, created_at FROM dayplans 
WHERE schedule_id = $1
`

func (q *Queries) GetDayplanForSchedules(ctx context.Context, scheduleID int32) ([]Dayplan, error) {
	rows, err := q.db.QueryContext(ctx, getDayplanForSchedules, scheduleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Dayplan{}
	for rows.Next() {
		var i Dayplan
		if err := rows.Scan(
			&i.ID,
			&i.StartTime,
			&i.EndTime,
			&i.ScheduleID,
			&i.Title,
			&i.Description,
			&i.PerformerName,
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

const updateDayPlan = `-- name: UpdateDayPlan :one
UPDATE dayplans SET
 start_time = CASE WHEN $1::boolean
        THEN $2::timestamp ELSE start_time END, 
 end_time = CASE WHEN $3::boolean
        THEN $4::timestamp ELSE end_time END,
 title = CASE WHEN $5::boolean
        THEN $6::text ELSE title END,
 description = CASE WHEN $7::boolean
        THEN $8::text ELSE description END,
 performer_name = CASE WHEN $9::boolean
        THEN $10::text ELSE performer_name END
WHERE id = $11 RETURNING id, start_time, end_time, schedule_id, title, description, performer_name, created_at
`

type UpdateDayPlanParams struct {
	StartTimeToUpdate     bool      `json:"start_time_to_update"`
	StartTime             time.Time `json:"start_time"`
	EndTimeToUpdate       bool      `json:"end_time_to_update"`
	EndTime               time.Time `json:"end_time"`
	TitleToUpdate         bool      `json:"title_to_update"`
	Title                 string    `json:"title"`
	DescriptionToUpdate   bool      `json:"description_to_update"`
	Description           string    `json:"description"`
	PerformerNameToUpdate bool      `json:"performer_name_to_update"`
	PerformerName         string    `json:"performer_name"`
	ID                    int32     `json:"id"`
}

func (q *Queries) UpdateDayPlan(ctx context.Context, arg UpdateDayPlanParams) (Dayplan, error) {
	row := q.db.QueryRowContext(ctx, updateDayPlan,
		arg.StartTimeToUpdate,
		arg.StartTime,
		arg.EndTimeToUpdate,
		arg.EndTime,
		arg.TitleToUpdate,
		arg.Title,
		arg.DescriptionToUpdate,
		arg.Description,
		arg.PerformerNameToUpdate,
		arg.PerformerName,
		arg.ID,
	)
	var i Dayplan
	err := row.Scan(
		&i.ID,
		&i.StartTime,
		&i.EndTime,
		&i.ScheduleID,
		&i.Title,
		&i.Description,
		&i.PerformerName,
		&i.CreatedAt,
	)
	return i, err
}
