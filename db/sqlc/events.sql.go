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
    status
) VALUES
  ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id, title, description, banner_image, start_date, end_date, venue, type, user_id, category, ticket_id, recurring, status, created_at
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
		&i.TicketID,
		&i.Recurring,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const createFavoriteEvent = `-- name: CreateFavoriteEvent :one
INSERT INTO events_favorites (
    event_id,
    user_id
) VALUES
  ($1, $2) RETURNING id, event_id, user_id, created_at
`

type CreateFavoriteEventParams struct {
	EventID int32 `json:"event_id"`
	UserID  int32 `json:"user_id"`
}

func (q *Queries) CreateFavoriteEvent(ctx context.Context, arg CreateFavoriteEventParams) (EventsFavorite, error) {
	row := q.db.QueryRowContext(ctx, createFavoriteEvent, arg.EventID, arg.UserID)
	var i EventsFavorite
	err := row.Scan(
		&i.ID,
		&i.EventID,
		&i.UserID,
		&i.CreatedAt,
	)
	return i, err
}

const createViewEvent = `-- name: CreateViewEvent :one
INSERT INTO events_views (
    event_id,
    user_id
) VALUES
  ($1, $2) RETURNING id, event_id, user_id, created_at
`

type CreateViewEventParams struct {
	EventID int32 `json:"event_id"`
	UserID  int32 `json:"user_id"`
}

func (q *Queries) CreateViewEvent(ctx context.Context, arg CreateViewEventParams) (EventsView, error) {
	row := q.db.QueryRowContext(ctx, createViewEvent, arg.EventID, arg.UserID)
	var i EventsView
	err := row.Scan(
		&i.ID,
		&i.EventID,
		&i.UserID,
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
SELECT id, title, description, banner_image, start_date, end_date, venue, type, user_id, category, ticket_id, recurring, status, created_at FROM events
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
		&i.TicketID,
		&i.Recurring,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const getEvents = `-- name: GetEvents :many
SELECT id, title, description, banner_image, start_date, end_date, venue, type, user_id, category, ticket_id, recurring, status, created_at FROM events e
WHERE e.status = $1 AND
 e.end_date >= CURRENT_DATE
ORDER BY e.id desc
LIMIT $2
OFFSET $3 ROWS
`

type GetEventsParams struct {
	Status int32 `json:"status"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetEvents(ctx context.Context, arg GetEventsParams) ([]Event, error) {
	rows, err := q.db.QueryContext(ctx, getEvents, arg.Status, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Event{}
	for rows.Next() {
		var i Event
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
			&i.TicketID,
			&i.Recurring,
			&i.Status,
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

const getEventsByLocation = `-- name: GetEventsByLocation :many
SELECT v.id, name, address, postal_code, city, province, country_code, venue_owner, v.banner_image, rating, longitude, latitude, v.status, v.created_at, e.id, title, description, e.banner_image, start_date, end_date, venue, type, user_id, category, ticket_id, recurring, e.status, e.created_at, point($1,$2) <@>  (point(v.longitude, v.latitude)::point) as distance
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
	ID            int32           `json:"id"`
	Name          string          `json:"name"`
	Address       sql.NullString  `json:"address"`
	PostalCode    sql.NullString  `json:"postal_code"`
	City          sql.NullString  `json:"city"`
	Province      sql.NullString  `json:"province"`
	CountryCode   sql.NullString  `json:"country_code"`
	VenueOwner    int32           `json:"venue_owner"`
	BannerImage   sql.NullString  `json:"banner_image"`
	Rating        sql.NullFloat64 `json:"rating"`
	Longitude     sql.NullFloat64 `json:"longitude"`
	Latitude      sql.NullFloat64 `json:"latitude"`
	Status        int32           `json:"status"`
	CreatedAt     sql.NullTime    `json:"created_at"`
	ID_2          int32           `json:"id_2"`
	Title         string          `json:"title"`
	Description   string          `json:"description"`
	BannerImage_2 string          `json:"banner_image_2"`
	StartDate     time.Time       `json:"start_date"`
	EndDate       time.Time       `json:"end_date"`
	Venue         int32           `json:"venue"`
	Type          int32           `json:"type"`
	UserID        int32           `json:"user_id"`
	Category      int32           `json:"category"`
	TicketID      sql.NullInt32   `json:"ticket_id"`
	Recurring     sql.NullBool    `json:"recurring"`
	Status_2      int32           `json:"status_2"`
	CreatedAt_2   sql.NullTime    `json:"created_at_2"`
	Distance      interface{}     `json:"distance"`
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
			&i.VenueOwner,
			&i.BannerImage,
			&i.Rating,
			&i.Longitude,
			&i.Latitude,
			&i.Status,
			&i.CreatedAt,
			&i.ID_2,
			&i.Title,
			&i.Description,
			&i.BannerImage_2,
			&i.StartDate,
			&i.EndDate,
			&i.Venue,
			&i.Type,
			&i.UserID,
			&i.Category,
			&i.TicketID,
			&i.Recurring,
			&i.Status_2,
			&i.CreatedAt_2,
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

const getEventsFilter = `-- name: GetEventsFilter :many
SELECT id, title, description, banner_image, start_date, end_date, venue, type, user_id, category, ticket_id, recurring, status, created_at FROM events e
WHERE e.status = $1 AND
e.category = $2 AND
e.end_date >= CURRENT_DATE
ORDER BY e.id desc
LIMIT $3
OFFSET $4 ROWS
`

type GetEventsFilterParams struct {
	Status   int32 `json:"status"`
	Category int32 `json:"category"`
	Limit    int32 `json:"limit"`
	Offset   int32 `json:"offset"`
}

func (q *Queries) GetEventsFilter(ctx context.Context, arg GetEventsFilterParams) ([]Event, error) {
	rows, err := q.db.QueryContext(ctx, getEventsFilter,
		arg.Status,
		arg.Category,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Event{}
	for rows.Next() {
		var i Event
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
			&i.TicketID,
			&i.Recurring,
			&i.Status,
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

const getFavoriteEvents = `-- name: GetFavoriteEvents :many
SELECT id, event_id, user_id, created_at FROM events_favorites 
where user_id = $1
ORDER BY id desc
`

func (q *Queries) GetFavoriteEvents(ctx context.Context, userID int32) ([]EventsFavorite, error) {
	rows, err := q.db.QueryContext(ctx, getFavoriteEvents, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []EventsFavorite{}
	for rows.Next() {
		var i EventsFavorite
		if err := rows.Scan(
			&i.ID,
			&i.EventID,
			&i.UserID,
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

const getViewedEvents = `-- name: GetViewedEvents :many
SELECT id, event_id, user_id, created_at FROM events_views 
where user_id = $1
ORDER BY id desc
`

func (q *Queries) GetViewedEvents(ctx context.Context, userID int32) ([]EventsView, error) {
	rows, err := q.db.QueryContext(ctx, getViewedEvents, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []EventsView{}
	for rows.Next() {
		var i EventsView
		if err := rows.Scan(
			&i.ID,
			&i.EventID,
			&i.UserID,
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

const updateEvent = `-- name: UpdateEvent :one
UPDATE events SET
    title = CASE WHEN $1::boolean
        THEN $2::text ELSE title END, 
    description = CASE WHEN $3::boolean
        THEN $4::text ELSE description END,
    banner_image = CASE WHEN $5::boolean
        THEN $6::text ELSE banner_image END,
    start_date = CASE WHEN $7::boolean
        THEN $8::timestamptz ELSE start_date END,
    end_date =CASE WHEN $9::boolean
        THEN $10::timestamptz ELSE end_date END,
    venue = CASE WHEN $11::boolean
        THEN $12::INTEGER ELSE venue END,
    type = CASE WHEN $13::boolean
        THEN $14::INTEGER ELSE type END,
    user_id = CASE WHEN $15::boolean
        THEN $16::INTEGER ELSE user_id END,
    category = CASE WHEN $17::boolean
        THEN $18::INTEGER ELSE category END,
    status = CASE WHEN $19::boolean
        THEN $20::INTEGER ELSE status END
    WHERE id= $21 RETURNING id, title, description, banner_image, start_date, end_date, venue, type, user_id, category, ticket_id, recurring, status, created_at
`

type UpdateEventParams struct {
	TitleToUpdate       bool      `json:"title_to_update"`
	Title               string    `json:"title"`
	DescriptionToUpdate bool      `json:"description_to_update"`
	Description         string    `json:"description"`
	BannerImageToUpdate bool      `json:"banner_image_to_update"`
	BannerImage         string    `json:"banner_image"`
	StartDateToUpdate   bool      `json:"start_date_to_update"`
	StartDate           time.Time `json:"start_date"`
	EndDateToUpdate     bool      `json:"end_date_to_update"`
	EndDate             time.Time `json:"end_date"`
	VenueToUpdate       bool      `json:"venue_to_update"`
	Venue               int32     `json:"venue"`
	TypeToUpdate        bool      `json:"type_to_update"`
	Type                int32     `json:"type"`
	UserIDToUpdate      bool      `json:"user_id_to_update"`
	UserID              int32     `json:"user_id"`
	CategoryToUpdate    bool      `json:"category_to_update"`
	Category            int32     `json:"category"`
	StatusToUpdate      bool      `json:"status_to_update"`
	Status              int32     `json:"status"`
	ID                  int32     `json:"id"`
}

func (q *Queries) UpdateEvent(ctx context.Context, arg UpdateEventParams) (Event, error) {
	row := q.db.QueryRowContext(ctx, updateEvent,
		arg.TitleToUpdate,
		arg.Title,
		arg.DescriptionToUpdate,
		arg.Description,
		arg.BannerImageToUpdate,
		arg.BannerImage,
		arg.StartDateToUpdate,
		arg.StartDate,
		arg.EndDateToUpdate,
		arg.EndDate,
		arg.VenueToUpdate,
		arg.Venue,
		arg.TypeToUpdate,
		arg.Type,
		arg.UserIDToUpdate,
		arg.UserID,
		arg.CategoryToUpdate,
		arg.Category,
		arg.StatusToUpdate,
		arg.Status,
		arg.ID,
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
		&i.TicketID,
		&i.Recurring,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
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
