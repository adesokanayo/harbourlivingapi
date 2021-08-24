package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateVenue(t *testing.T) {

	arg := CreateVenueParams{
		Name:        "Test Venue",
		Address:     "adekunle ",
		City:        "lagos",
		Province:    "AB",
		CountryCode: "CAN",
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

}

func CreateRandomVenue() Venue {
	arg := CreateVenueParams{
		Name:        "Test Venue",
		Address:     "lagos",
		City:        "lagos",
		Province:    "AB",
		CountryCode: "CAN",
	}
	venue, err := testQueries.CreateVenue(context.Background(), arg)
	if err != nil {

	}
	return venue
}
