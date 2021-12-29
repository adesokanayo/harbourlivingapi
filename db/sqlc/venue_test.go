package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateVenue(t *testing.T) {

	arg := CreateVenueParams{
		Name: "Test Venue",
		Address: sql.NullString{
			String: "adekunle",
			Valid:  true,
		},
		City:        sql.NullString{String: "Calgary", Valid: true},
		Province:    sql.NullString{String: "AB", Valid: true},
		CountryCode: sql.NullString{String: "CAN", Valid: true},
		Status: 1,
	}
	venue, err := testQueries.CreateVenue(context.Background(), arg)

	require.Empty(t, err)
	require.NotEmpty(t, venue)
	require.Equal(t, venue.City, arg.City)
	require.Equal(t, venue.Address, arg.Address)
	require.Equal(t, venue.Province, arg.Province)
	require.Equal(t, venue.Name, arg.Name)

}

func TestGetVenue(t *testing.T) {
	randomVenue := CreateRandomVenue()
	venue, err := testQueries.GetVenue(context.Background(), randomVenue.ID)

	require.Empty(t, err)
	require.Equal(t, venue.City, randomVenue.City)
	require.Equal(t, venue.Address, randomVenue.Address)
	require.Equal(t, venue.Province, randomVenue.Province)
	require.Equal(t, venue.Name, randomVenue.Name)
	require.Equal(t, venue.Status, randomVenue.Status)

}

func CreateRandomVenue() Venue {
	arg := CreateVenueParams{
		Name: "Test Venue",
		Address: sql.NullString{
			String: "adekunle",
			Valid:  true,
		},
		City:        sql.NullString{String: "Calgary", Valid: true},
		Province:    sql.NullString{String: "AB", Valid: true},
		CountryCode: sql.NullString{String: "CAN", Valid: true},
		Status:      1,
	}

	venue, err := testQueries.CreateVenue(context.Background(), arg)
	if err != nil {

	}
	return venue
}
