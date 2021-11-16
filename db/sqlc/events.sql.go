// Code generated by sqlc. DO NOT EDIT.
// source: events.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createEvent = `-- name: CreateEvent :one
INSERT INTO events (
    title,
    description,
    banner_image,
    start_date,
    end_date,
    venue,
    type,
    user_id,
    category,
    subcategory,
    status
) VALUES
  ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id, title, description, banner_image, start_date, end_date, venue, type, user_id, category, subcategory, ticket_id, recurring, status, created_at
`

type CreateEventParams struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	BannerImage string    `json:"banner_image"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Venue       int32     `json:"venue"`
	Type        int32     `json:"type"`
	UserID      int32     `json:"user_id"`
	Category    int32     `json:"category"`
	Subcategory int32     `json:"subcategory"`
	Status      int32     `json:"status"`
}

func (q *Queries) CreateEvent(ctx context.Context, arg CreateEventParams) (Event, error) {
	row := q.db.QueryRowContext(ctx, createEvent,
		arg.Title,
		arg.Description,
		arg.BannerImage,
		arg.StartDate,
		arg.EndDate,
		arg.Venue,
		arg.Type,
		arg.UserID,
		arg.Category,
		arg.Subcategory,
		arg.Status,
	)
	var i Event
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.BannerImage,
		&i.StartDate,
		&i.EndDate,
		&i.Venue,
		&i.Type,
		&i.UserID,
		&i.Category,
		&i.Subcategory,
		&i.TicketID,
		&i.Recurring,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const deleteEvent = `-- name: DeleteEvent :exec
DELETE FROM events
WHERE id = $1
`

func (q *Queries) DeleteEvent(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteEvent, id)
	return err
}

const getEvent = `-- name: GetEvent :one
SELECT id, title, description, banner_image, start_date, end_date, venue, type, user_id, category, subcategory, ticket_id, recurring, status, created_at FROM events
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetEvent(ctx context.Context, id int32) (Event, error) {
	row := q.db.QueryRowContext(ctx, getEvent, id)
	var i Event
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.BannerImage,
		&i.StartDate,
		&i.EndDate,
		&i.Venue,
		&i.Type,
		&i.UserID,
		&i.Category,
		&i.Subcategory,
		&i.TicketID,
		&i.Recurring,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const getEvents = `-- name: GetEvents :many
SELECT e.id, title, description, banner_image, start_date, end_date, venue, type, user_id, category, subcategory, ticket_id, recurring, status, created_at, v.id, name, address, postal_code, city, province, country_code, url, virtual, rating, longitude, latitude FROM events e
inner join venues v on  e.venue = v.id
WHERE category = $1
and subcategory =$2
and e.status =$3
ORDER BY e.id desc
LIMIT $4
OFFSET $5
`

type GetEventsParams struct {
	Category    int32 `json:"category"`
	Subcategory int32 `json:"subcategory"`
	Status      int32 `json:"status"`
	Limit       int32 `json:"limit"`
	Offset      int32 `json:"offset"`
}

type GetEventsRow struct {
	ID          int32           `json:"id"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	BannerImage string          `json:"banner_image"`
	StartDate   time.Time       `json:"start_date"`
	EndDate     time.Time       `json:"end_date"`
	Venue       int32           `json:"venue"`
	Type        int32           `json:"type"`
	UserID      int32           `json:"user_id"`
	Category    int32           `json:"category"`
	Subcategory int32           `json:"subcategory"`
	TicketID    sql.NullInt32   `json:"ticket_id"`
	Recurring   sql.NullBool    `json:"recurring"`
	Status      int32           `json:"status"`
	CreatedAt   sql.NullTime    `json:"created_at"`
	ID_2        int32           `json:"id_2"`
	Name        string          `json:"name"`
	Address     sql.NullString  `json:"address"`
	PostalCode  sql.NullString  `json:"postal_code"`
	City        sql.NullString  `json:"city"`
	Province    sql.NullString  `json:"province"`
	CountryCode sql.NullString  `json:"country_code"`
	Url         sql.NullString  `json:"url"`
	Virtual     bool            `json:"virtual"`
	Rating      sql.NullFloat64 `json:"rating"`
	Longitude   sql.NullFloat64 `json:"longitude"`
	Latitude    sql.NullFloat64 `json:"latitude"`
}

func (q *Queries) GetEvents(ctx context.Context, arg GetEventsParams) ([]GetEventsRow, error) {
	rows, err := q.db.QueryContext(ctx, getEvents,
		arg.Category,
		arg.Subcategory,
		arg.Status,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetEventsRow{}
	for rows.Next() {
		var i GetEventsRow
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.BannerImage,
			&i.StartDate,
			&i.EndDate,
			&i.Venue,
			&i.Type,
			&i.UserID,
			&i.Category,
			&i.Subcategory,
			&i.TicketID,
			&i.Recurring,
			&i.Status,
			&i.CreatedAt,
			&i.ID_2,
			&i.Name,
			&i.Address,
			&i.PostalCode,
			&i.City,
			&i.Province,
			&i.CountryCode,
			&i.Url,
			&i.Virtual,
			&i.Rating,
			&i.Longitude,
			&i.Latitude,
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

const getEventsByLocation = `-- name: GetEventsByLocation :many
SELECT v.id, name, address, postal_code, city, province, country_code, url, virtual, rating, longitude, latitude, e.id, title, description, banner_image, start_date, end_date, venue, type, user_id, category, subcategory, ticket_id, recurring, status, created_at, point($1,$2) <@>  (point(v.longitude, v.latitude)::point) as distance
FROM venues v, events e
WHERE (point($1,$2) <@> point(longitude, latitude)) < $3  
ORDER BY distance desc
`

type GetEventsByLocationParams struct {
	Point     float64         `json:"point"`
	Point_2   float64         `json:"point_2"`
	Longitude sql.NullFloat64 `json:"longitude"`
}

type GetEventsByLocationRow struct {
	ID          int32           `json:"id"`
	Name        string          `json:"name"`
	Address     sql.NullString  `json:"address"`
	PostalCode  sql.NullString  `json:"postal_code"`
	City        sql.NullString  `json:"city"`
	Province    sql.NullString  `json:"province"`
	CountryCode sql.NullString  `json:"country_code"`
	Url         sql.NullString  `json:"url"`
	Virtual     bool            `json:"virtual"`
	Rating      sql.NullFloat64 `json:"rating"`
	Longitude   sql.NullFloat64 `json:"longitude"`
	Latitude    sql.NullFloat64 `json:"latitude"`
	ID_2        int32           `json:"id_2"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	BannerImage string          `json:"banner_image"`
	StartDate   time.Time       `json:"start_date"`
	EndDate     time.Time       `json:"end_date"`
	Venue       int32           `json:"venue"`
	Type        int32           `json:"type"`
	UserID      int32           `json:"user_id"`
	Category    int32           `json:"category"`
	Subcategory int32           `json:"subcategory"`
	TicketID    sql.NullInt32   `json:"ticket_id"`
	Recurring   sql.NullBool    `json:"recurring"`
	Status      int32           `json:"status"`
	CreatedAt   sql.NullTime    `json:"created_at"`
	Distance    interface{}     `json:"distance"`
}

func (q *Queries) GetEventsByLocation(ctx context.Context, arg GetEventsByLocationParams) ([]GetEventsByLocationRow, error) {
	rows, err := q.db.QueryContext(ctx, getEventsByLocation, arg.Point, arg.Point_2, arg.Longitude)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetEventsByLocationRow{}
	for rows.Next() {
		var i GetEventsByLocationRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Address,
			&i.PostalCode,
			&i.City,
			&i.Province,
			&i.CountryCode,
			&i.Url,
			&i.Virtual,
			&i.Rating,
			&i.Longitude,
			&i.Latitude,
			&i.ID_2,
			&i.Title,
			&i.Description,
			&i.BannerImage,
			&i.StartDate,
			&i.EndDate,
			&i.Venue,
			&i.Type,
			&i.UserID,
			&i.Category,
			&i.Subcategory,
			&i.TicketID,
			&i.Recurring,
			&i.Status,
			&i.CreatedAt,
			&i.Distance,
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

const updateEventStatus = `-- name: UpdateEventStatus :one
UPDATE events
set status = $1
where id = $2 RETURNING events.Id, events.status
`

type UpdateEventStatusParams struct {
	Status int32 `json:"status"`
	ID     int32 `json:"id"`
}

type UpdateEventStatusRow struct {
	ID     int32 `json:"id"`
	Status int32 `json:"status"`
}

func (q *Queries) UpdateEventStatus(ctx context.Context, arg UpdateEventStatusParams) (UpdateEventStatusRow, error) {
	row := q.db.QueryRowContext(ctx, updateEventStatus, arg.Status, arg.ID)
	var i UpdateEventStatusRow
	err := row.Scan(&i.ID, &i.Status)
	return i, err
}
