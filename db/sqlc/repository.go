package db

import (
	"context"
	"database/sql"
)

type Repository interface {
	CreateEvent(ctx context.Context, arg CreateEventParams) (Event, error)
	DeleteEvents(ctx context.Context, id int32) error
	GetEvents(ctx context.Context, arg GetEventsParams) ([]Event, error)
	GetEvent(ctx context.Context, id int32) (Event, error)

	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteUser(ctx context.Context, id int32) error
	GetAllUsers(ctx context.Context) ([]User, error)
	GetUser(ctx context.Context, id int32) (User, error)
	GetUsername(ctx context.Context, username string) (User, error)

	CreateVenue(ctx context.Context, arg CreateVenueParams) (Venue, error)
	DeleteVenue(ctx context.Context, id int32) error
	GetAllVenues(ctx context.Context) ([]Venue, error)
	GetVenue(ctx context.Context, id int32) (Venue, error)
}

type repoSvc struct {
	*Queries
	db *sql.DB
}
