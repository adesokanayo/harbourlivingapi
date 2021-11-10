// Code generated by sqlc. DO NOT EDIT.
// source: video.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createVideo = `-- name: CreateVideo :one
INSERT INTO Videos ( name, url  )
VALUES ($1,$2) RETURNING id, name, url
`

type CreateVideoParams struct {
	Name sql.NullString `json:"name"`
	Url  string         `json:"url"`
}

func (q *Queries) CreateVideo(ctx context.Context, arg CreateVideoParams) (Video, error) {
	row := q.db.QueryRowContext(ctx, createVideo, arg.Name, arg.Url)
	var i Video
	err := row.Scan(&i.ID, &i.Name, &i.Url)
	return i, err
}

const getAllVideos = `-- name: GetAllVideos :many
SELECT id, name, url FROM Videos
`

func (q *Queries) GetAllVideos(ctx context.Context) ([]Video, error) {
	rows, err := q.db.QueryContext(ctx, getAllVideos)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Video{}
	for rows.Next() {
		var i Video
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

const getVideo = `-- name: GetVideo :one
SELECT id, name, url FROM Videos
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetVideo(ctx context.Context, id int32) (Video, error) {
	row := q.db.QueryRowContext(ctx, getVideo, id)
	var i Video
	err := row.Scan(&i.ID, &i.Name, &i.Url)
	return i, err
}

const getVideosByEvent = `-- name: GetVideosByEvent :many
SELECT a.id, name, url, b.id, event_id, video_id, created_at FROM videos a, events_videos b
WHERE event_id = $1 and a.event_id = b.event_id
`

type GetVideosByEventRow struct {
	ID        int32          `json:"id"`
	Name      sql.NullString `json:"name"`
	Url       string         `json:"url"`
	ID_2      int32          `json:"id_2"`
	EventID   int32          `json:"event_id"`
	VideoID   int32          `json:"video_id"`
	CreatedAt time.Time      `json:"created_at"`
}

func (q *Queries) GetVideosByEvent(ctx context.Context, eventID int32) ([]GetVideosByEventRow, error) {
	rows, err := q.db.QueryContext(ctx, getVideosByEvent, eventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetVideosByEventRow{}
	for rows.Next() {
		var i GetVideosByEventRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Url,
			&i.ID_2,
			&i.EventID,
			&i.VideoID,
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

const linkVideoToEvent = `-- name: LinkVideoToEvent :exec
INSERT into events_videos (event_id, video_id)
VALUES($1,$2)
`

type LinkVideoToEventParams struct {
	EventID int32 `json:"event_id"`
	VideoID int32 `json:"video_id"`
}

func (q *Queries) LinkVideoToEvent(ctx context.Context, arg LinkVideoToEventParams) error {
	_, err := q.db.ExecContext(ctx, linkVideoToEvent, arg.EventID, arg.VideoID)
	return err
}
