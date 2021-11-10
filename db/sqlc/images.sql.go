// Code generated by sqlc. DO NOT EDIT.
// source: images.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createImage = `-- name: CreateImage :one
INSERT INTO images ( name, url  )
VALUES ($1,$2) RETURNING id, name, url
`

type CreateImageParams struct {
	Name sql.NullString `json:"name"`
	Url  string         `json:"url"`
}

func (q *Queries) CreateImage(ctx context.Context, arg CreateImageParams) (Image, error) {
	row := q.db.QueryRowContext(ctx, createImage, arg.Name, arg.Url)
	var i Image
	err := row.Scan(&i.ID, &i.Name, &i.Url)
	return i, err
}

const getAllImages = `-- name: GetAllImages :many
SELECT id, name, url FROM images
`

func (q *Queries) GetAllImages(ctx context.Context) ([]Image, error) {
	rows, err := q.db.QueryContext(ctx, getAllImages)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Image{}
	for rows.Next() {
		var i Image
		if err := rows.Scan(&i.ID, &i.Name, &i.Url); err != nil {
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

const getImage = `-- name: GetImage :one
SELECT id, name, url FROM images
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetImage(ctx context.Context, id int32) (Image, error) {
	row := q.db.QueryRowContext(ctx, getImage, id)
	var i Image
	err := row.Scan(&i.ID, &i.Name, &i.Url)
	return i, err
}

const getImagesByEvent = `-- name: GetImagesByEvent :many
SELECT a.id, name, url, b.id, event_id, image_id, created_at FROM images a, events_images b
WHERE event_id = $1 and a.event_id = b.event_id
`

type GetImagesByEventRow struct {
	ID        int32          `json:"id"`
	Name      sql.NullString `json:"name"`
	Url       string         `json:"url"`
	ID_2      int32          `json:"id_2"`
	EventID   int32          `json:"event_id"`
	ImageID   int32          `json:"image_id"`
	CreatedAt time.Time      `json:"created_at"`
}

func (q *Queries) GetImagesByEvent(ctx context.Context, eventID int32) ([]GetImagesByEventRow, error) {
	rows, err := q.db.QueryContext(ctx, getImagesByEvent, eventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetImagesByEventRow{}
	for rows.Next() {
		var i GetImagesByEventRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Url,
			&i.ID_2,
			&i.EventID,
			&i.ImageID,
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

const linkImageToEvent = `-- name: LinkImageToEvent :exec
INSERT into events_images (event_id, image_id)
VALUES($1,$2)
`

type LinkImageToEventParams struct {
	EventID int32 `json:"event_id"`
	ImageID int32 `json:"image_id"`
}

func (q *Queries) LinkImageToEvent(ctx context.Context, arg LinkImageToEventParams) error {
	_, err := q.db.ExecContext(ctx, linkImageToEvent, arg.EventID, arg.ImageID)
	return err
}
