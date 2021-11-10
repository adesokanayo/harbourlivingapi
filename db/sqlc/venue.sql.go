// Code generated by sqlc. DO NOT EDIT.
// source: venue.sql

package db

import (
	"context"
	"database/sql"
)

const createVenue = `-- name: CreateVenue :one
INSERT INTO venue (
    name,
    address,
    postal_code,
    city,
    province,
    country_code,
    url,
    virtual,
    longitude, 
    latitude
) VALUES
    ($1, $2, $3, $4, $5, $6,$7, $8,$9,$10) RETURNING id, name, address, postal_code, city, province, country_code, url, virtual, rating, longitude, latitude
`

type CreateVenueParams struct {
	Name        string          `json:"name"`
	Address     sql.NullString  `json:"address"`
	PostalCode  sql.NullString  `json:"postal_code"`
	City        sql.NullString  `json:"city"`
	Province    sql.NullString  `json:"province"`
	CountryCode sql.NullString  `json:"country_code"`
	Url         sql.NullString  `json:"url"`
	Virtual     bool            `json:"virtual"`
	Longitude   sql.NullFloat64 `json:"longitude"`
	Latitude    sql.NullFloat64 `json:"latitude"`
}

func (q *Queries) CreateVenue(ctx context.Context, arg CreateVenueParams) (Venue, error) {
	row := q.db.QueryRowContext(ctx, createVenue,
		arg.Name,
		arg.Address,
		arg.PostalCode,
		arg.City,
		arg.Province,
		arg.CountryCode,
		arg.Url,
		arg.Virtual,
		arg.Longitude,
		arg.Latitude,
	)
	var i Venue
	err := row.Scan(
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
	)
	return i, err
}

const createVirtualVenue = `-- name: CreateVirtualVenue :one
INSERT INTO venue (
    name,
    url,
    virtual
) VALUES
    ($1, $2, $3) RETURNING id, name, address, postal_code, city, province, country_code, url, virtual, rating, longitude, latitude
`

type CreateVirtualVenueParams struct {
	Name    string         `json:"name"`
	Url     sql.NullString `json:"url"`
	Virtual bool           `json:"virtual"`
}

func (q *Queries) CreateVirtualVenue(ctx context.Context, arg CreateVirtualVenueParams) (Venue, error) {
	row := q.db.QueryRowContext(ctx, createVirtualVenue, arg.Name, arg.Url, arg.Virtual)
	var i Venue
	err := row.Scan(
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
	)
	return i, err
}

const deleteVenue = `-- name: DeleteVenue :exec
DELETE FROM venue
WHERE id = $1
`

func (q *Queries) DeleteVenue(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteVenue, id)
	return err
}

const getAllVenues = `-- name: GetAllVenues :many
SELECT id, name, address, postal_code, city, province, country_code, url, virtual, rating, longitude, latitude FROM venue
ORDER  by id
`

func (q *Queries) GetAllVenues(ctx context.Context) ([]Venue, error) {
	rows, err := q.db.QueryContext(ctx, getAllVenues)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Venue{}
	for rows.Next() {
		var i Venue
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

const getVenue = `-- name: GetVenue :one
SELECT id, name, address, postal_code, city, province, country_code, url, virtual, rating, longitude, latitude FROM venue
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetVenue(ctx context.Context, id int32) (Venue, error) {
	row := q.db.QueryRowContext(ctx, getVenue, id)
	var i Venue
	err := row.Scan(
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
	)
	return i, err
}

const rateVenue = `-- name: RateVenue :exec
UPDATE venue SET rating = $1
where id = $2
`

type RateVenueParams struct {
	Rating sql.NullFloat64 `json:"rating"`
	ID     int32           `json:"id"`
}

func (q *Queries) RateVenue(ctx context.Context, arg RateVenueParams) error {
	_, err := q.db.ExecContext(ctx, rateVenue, arg.Rating, arg.ID)
	return err
}
