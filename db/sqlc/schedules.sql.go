// Code generated by sqlc. DO NOT EDIT.
// source: schedules.sql

package db

import (
	"context"
	"time"
)

const createSchedule = `-- name: CreateSchedule :one
INSERT INTO schedules (
    event_id,
    date,
    start_time,
    end_time
) VALUES
  ($1, $2, $3, $4 ) RETURNING id, event_id, date, start_time, end_time, created_at
`

type CreateScheduleParams struct {
	EventID   int32     `json:"event_id"`
	Date      time.Time `json:"date"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

func (q *Queries) CreateSchedule(ctx context.Context, arg CreateScheduleParams) (Schedule, error) {
	row := q.db.QueryRowContext(ctx, createSchedule,
		arg.EventID,
		arg.Date,
		arg.StartTime,
		arg.EndTime,
	)
	var i Schedule
	err := row.Scan(
		&i.ID,
		&i.EventID,
		&i.Date,
		&i.StartTime,
		&i.EndTime,
		&i.CreatedAt,
	)
	return i, err
}

const deleteSchedule = `-- name: DeleteSchedule :exec
DELETE FROM schedules
WHERE id = $1
`

func (q *Queries) DeleteSchedule(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteSchedule, id)
	return err
}

const getAllSchedule = `-- name: GetAllSchedule :many
SELECT id, event_id, date, start_time, end_time, created_at  FROM schedules
order by id desc
`

func (q *Queries) GetAllSchedule(ctx context.Context) ([]Schedule, error) {
	rows, err := q.db.QueryContext(ctx, getAllSchedule)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Schedule{}
	for rows.Next() {
		var i Schedule
		if err := rows.Scan(
			&i.ID,
			&i.EventID,
			&i.Date,
			&i.StartTime,
			&i.EndTime,
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

const getSchedule = `-- name: GetSchedule :one
SELECT id, event_id, date, start_time, end_time, created_at FROM schedules
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetSchedule(ctx context.Context, id int32) (Schedule, error) {
	row := q.db.QueryRowContext(ctx, getSchedule, id)
	var i Schedule
	err := row.Scan(
		&i.ID,
		&i.EventID,
		&i.Date,
		&i.StartTime,
		&i.EndTime,
		&i.CreatedAt,
	)
	return i, err
}

const getSchedulesForEvent = `-- name: GetSchedulesForEvent :many
SELECT id, event_id, date, start_time, end_time, created_at FROM schedules
WHERE event_id = $1
`

func (q *Queries) GetSchedulesForEvent(ctx context.Context, eventID int32) ([]Schedule, error) {
	rows, err := q.db.QueryContext(ctx, getSchedulesForEvent, eventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Schedule{}
	for rows.Next() {
		var i Schedule
		if err := rows.Scan(
			&i.ID,
			&i.EventID,
			&i.Date,
			&i.StartTime,
			&i.EndTime,
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

const updateSchedule = `-- name: UpdateSchedule :one
UPDATE schedules SET
 event_id = CASE WHEN $1::boolean
        THEN $2::int ELSE event_id END, 
 date = CASE WHEN $3::boolean
        THEN $4::timestamp ELSE date END,
 start_time = CASE WHEN $5::boolean
        THEN $6::timestamp ELSE start_time END,
 end_time = CASE WHEN $7::boolean
        THEN $8::timestamp ELSE end_time END
WHERE id = $9 RETURNING id, event_id, date, start_time, end_time, created_at
`

type UpdateScheduleParams struct {
	EventIDToUpdate   bool      `json:"event_id_to_update"`
	EventID           int32     `json:"event_id"`
	DateIDToUpdate    bool      `json:"date_id_to_update"`
	Date              time.Time `json:"date"`
	StartTimeToUpdate bool      `json:"start_time_to_update"`
	StartTime         time.Time `json:"start_time"`
	EndTimeToUpdate   bool      `json:"end_time_to_update"`
	EndTime           time.Time `json:"end_time"`
	ID                int32     `json:"id"`
}

func (q *Queries) UpdateSchedule(ctx context.Context, arg UpdateScheduleParams) (Schedule, error) {
	row := q.db.QueryRowContext(ctx, updateSchedule,
		arg.EventIDToUpdate,
		arg.EventID,
		arg.DateIDToUpdate,
		arg.Date,
		arg.StartTimeToUpdate,
		arg.StartTime,
		arg.EndTimeToUpdate,
		arg.EndTime,
		arg.ID,
	)
	var i Schedule
	err := row.Scan(
		&i.ID,
		&i.EventID,
		&i.Date,
		&i.StartTime,
		&i.EndTime,
		&i.CreatedAt,
	)
	return i, err
}
