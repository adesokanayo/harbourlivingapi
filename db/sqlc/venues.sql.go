// Code generated by sqlc. DO NOT EDIT.
// source: venues.sql

package db

import (
	"context"
	"database/sql"
)

const createVenue = `-- name: CreateVenue :one
INSERT INTO venues (
    name,
    address,
    postal_code,
    city,
    province,
    country_code,
    venue_owner,
    banner_image,
    longitude, 
    latitude,
    rating,
    status
) VALUES
    ($1, $2, $3, $4, $5, $6,$7, $8,$9,$10,$11,$12) RETURNING id, name, address, postal_code, city, province, country_code, venue_owner, banner_image, rating, longitude, latitude, status, created_at
`

type CreateVenueParams struct {
	Name        string          `json:"name"`
	Address     sql.NullString  `json:"address"`
	PostalCode  sql.NullString  `json:"postal_code"`
	City        sql.NullString  `json:"city"`
	Province    sql.NullString  `json:"province"`
	CountryCode sql.NullString  `json:"country_code"`
	VenueOwner  int32           `json:"venue_owner"`
	BannerImage sql.NullString  `json:"banner_image"`
	Longitude   sql.NullFloat64 `json:"longitude"`
	Latitude    sql.NullFloat64 `json:"latitude"`
	Rating      sql.NullFloat64 `json:"rating"`
	Status      int32           `json:"status"`
}

func (q *Queries) CreateVenue(ctx context.Context, arg CreateVenueParams) (Venue, error) {
	row := q.db.QueryRowContext(ctx, createVenue,
		arg.Name,
		arg.Address,
		arg.PostalCode,
		arg.City,
		arg.Province,
		arg.CountryCode,
		arg.VenueOwner,
		arg.BannerImage,
		arg.Longitude,
		arg.Latitude,
		arg.Rating,
		arg.Status,
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
		&i.VenueOwner,
		&i.BannerImage,
		&i.Rating,
		&i.Longitude,
		&i.Latitude,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const createVenueFavorite = `-- name: CreateVenueFavorite :one
INSERT INTO venues_favorites (
    venue_id,
    user_id
) VALUES
  ($1, $2) RETURNING id, venue_id, user_id, created_at
`

type CreateVenueFavoriteParams struct {
	VenueID int32 `json:"venue_id"`
	UserID  int32 `json:"user_id"`
}

func (q *Queries) CreateVenueFavorite(ctx context.Context, arg CreateVenueFavoriteParams) (VenuesFavorite, error) {
	row := q.db.QueryRowContext(ctx, createVenueFavorite, arg.VenueID, arg.UserID)
	var i VenuesFavorite
	err := row.Scan(
		&i.ID,
		&i.VenueID,
		&i.UserID,
		&i.CreatedAt,
	)
	return i, err
}

const deleteFavoriteVenue = `-- name: DeleteFavoriteVenue :exec
DELETE FROM venues_favorites 
WHERE venue_id = $1 
AND user_id = $2
`

type DeleteFavoriteVenueParams struct {
	VenueID int32 `json:"venue_id"`
	UserID  int32 `json:"user_id"`
}

func (q *Queries) DeleteFavoriteVenue(ctx context.Context, arg DeleteFavoriteVenueParams) error {
	_, err := q.db.ExecContext(ctx, deleteFavoriteVenue, arg.VenueID, arg.UserID)
	return err
}

const deleteVenue = `-- name: DeleteVenue :exec
DELETE FROM venues
WHERE id = $1
`

func (q *Queries) DeleteVenue(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteVenue, id)
	return err
}

const getAllVenues = `-- name: GetAllVenues :many
SELECT id, name, address, postal_code, city, province, country_code, venue_owner, banner_image, rating, longitude, latitude, status, created_at FROM venues
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
			&i.VenueOwner,
			&i.BannerImage,
			&i.Rating,
			&i.Longitude,
			&i.Latitude,
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

const getFavoriteVenues = `-- name: GetFavoriteVenues :many
SELECT id, venue_id, user_id, created_at FROM venues_favorites 
where user_id = $1
ORDER BY id desc
`

func (q *Queries) GetFavoriteVenues(ctx context.Context, userID int32) ([]VenuesFavorite, error) {
	rows, err := q.db.QueryContext(ctx, getFavoriteVenues, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []VenuesFavorite{}
	for rows.Next() {
		var i VenuesFavorite
		if err := rows.Scan(
			&i.ID,
			&i.VenueID,
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

const getVenue = `-- name: GetVenue :one
SELECT id, name, address, postal_code, city, province, country_code, venue_owner, banner_image, rating, longitude, latitude, status, created_at FROM venues
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
		&i.VenueOwner,
		&i.BannerImage,
		&i.Rating,
		&i.Longitude,
		&i.Latitude,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const rateVenue = `-- name: RateVenue :exec
UPDATE venues SET rating = $1
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

const updateVenue = `-- name: UpdateVenue :one
UPDATE venues SET
    name = CASE WHEN $1::boolean
        THEN $2::text ELSE name END, 
    address = CASE WHEN $3::boolean
        THEN $4::text ELSE address END,
    postal_code = CASE WHEN $5::boolean
        THEN $6::text ELSE postal_code END,
    city = CASE WHEN $7::boolean
        THEN $8::text ELSE city END,
    province = CASE WHEN $9::boolean
        THEN $10::text ELSE province END,
    country_code = CASE WHEN $11::boolean
        THEN $12::text ELSE country_code END,
    banner_image = CASE WHEN $13::boolean
        THEN $14::text ELSE banner_image END,
    longitude = CASE WHEN $15::boolean
        THEN $16::float ELSE longitude END,
    latitude = CASE WHEN $17::boolean
        THEN $18::float ELSE latitude END,
    rating = CASE WHEN $19::boolean
        THEN $20::int ELSE rating END,
    status = CASE WHEN $21::boolean
        THEN $22::INTEGER ELSE status END
    WHERE id= $23 RETURNING id, name, address, postal_code, city, province, country_code, venue_owner, banner_image, rating, longitude, latitude, status, created_at
`

type UpdateVenueParams struct {
	NameToUpdate        bool    `json:"name_to_update"`
	Name                string  `json:"name"`
	AddressToUpdate     bool    `json:"address_to_update"`
	Address             string  `json:"address"`
	PostalCodeToUpdate  bool    `json:"postal_code_to_update"`
	PostalCode          string  `json:"postal_code"`
	CityToUpdate        bool    `json:"city_to_update"`
	City                string  `json:"city"`
	ProvinceToUpdate    bool    `json:"province_to_update"`
	Province            string  `json:"province"`
	CountryToUpdate     bool    `json:"country_to_update"`
	CountryCode         string  `json:"country_code"`
	BannerImageToUpdate bool    `json:"banner_image_to_update"`
	BannerImage         string  `json:"banner_image"`
	LongitudeToUpdate   bool    `json:"longitude_to_update"`
	Longitude           float64 `json:"longitude"`
	LatitudeToUpdate    bool    `json:"latitude_to_update"`
	Latitude            float64 `json:"latitude"`
	RatingToUpdate      bool    `json:"rating_to_update"`
	Rating              int32   `json:"rating"`
	StatusToUpdate      bool    `json:"status_to_update"`
	Status              int32   `json:"status"`
	ID                  int32   `json:"id"`
}

func (q *Queries) UpdateVenue(ctx context.Context, arg UpdateVenueParams) (Venue, error) {
	row := q.db.QueryRowContext(ctx, updateVenue,
		arg.NameToUpdate,
		arg.Name,
		arg.AddressToUpdate,
		arg.Address,
		arg.PostalCodeToUpdate,
		arg.PostalCode,
		arg.CityToUpdate,
		arg.City,
		arg.ProvinceToUpdate,
		arg.Province,
		arg.CountryToUpdate,
		arg.CountryCode,
		arg.BannerImageToUpdate,
		arg.BannerImage,
		arg.LongitudeToUpdate,
		arg.Longitude,
		arg.LatitudeToUpdate,
		arg.Latitude,
		arg.RatingToUpdate,
		arg.Rating,
		arg.StatusToUpdate,
		arg.Status,
		arg.ID,
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
		&i.VenueOwner,
		&i.BannerImage,
		&i.Rating,
		&i.Longitude,
		&i.Latitude,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}
