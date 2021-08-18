package db

import (
	"context"
	"database/sql"
	"github.com/BigListRyRy/harbourlivingapi/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomEvent(t *testing.T) Event {
	user := createRandomUser(t)
	arg := CreateEventParams{
		Title:       util.RandomString(7),
		Description: util.RandomString(20),
		BannerImage: util.RandomString(10),
		StartDate:   util.RandomTime(1977, 2023),
		EndDate:     util.RandomTime(1977, 2023),
		Venue:       1,
		UserID:      user.ID,
		Type:        1,
		Category:    1,
		Subcategory: 1,
	}

	event, err := testQueries.CreateEvent(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, event)

	require.Equal(t, arg.Title, event.Title)
	require.Equal(t, arg.Description, event.Description)
	require.Equal(t, arg.BannerImage, event.BannerImage)
	require.Equal(t, arg.Status, event.Status)
	require.Equal(t, arg.Venue, event.Venue)
	require.Equal(t, arg.Type, event.Type)
	require.Equal(t, arg.UserID, event.UserID)
	require.Equal(t, arg.Category, event.Category)
	require.Equal(t, arg.Subcategory, event.Subcategory)
	require.Equal(t, arg.Image1, event.Image1)
	require.Equal(t, arg.Image2, event.Image2)
	require.NotZero(t, event.ID)
	require.NotZero(t, event.StartDate)
	require.NotZero(t, event.EndDate)
	require.NotZero(t, event.CreatedAt)

	return event
}

func TestCreateEvent(t *testing.T) {
	createRandomEvent(t)
}

func TestGetEvent(t *testing.T) {
	event1 := createRandomEvent(t)
	event2, err := testQueries.GetEvent(context.Background(), event1.ID)
	require.NoError(t, err)
	require.Equal(t, event1.Title, event2.Title)
	require.Equal(t, event1.Description, event2.Description)
	require.Equal(t, event1.BannerImage, event2.BannerImage)
	require.Equal(t, event1.Image2, event2.Image2)
}

func TestGetAllEvents(t *testing.T) {

	for i := 0; i < 10; i++ {
		createRandomEvent(t)
	}

	arg := GetAllEventsParams{
		Limit:       1,
		Offset:      1,
		Category:    1,
		Subcategory: 1,
	}
	events, err := testQueries.GetAllEvents(context.Background(), arg)
	require.NoError(t, err)
	for _, event := range events {
		require.NotEmpty(t, event)
	}
}

func TestDeleteEvent(t *testing.T) {
	event1 := createRandomEvent(t)
	err := testQueries.DeleteEvents(context.Background(), event1.ID)
	require.NoError(t, err)
	event2, err := testQueries.GetEvent(context.Background(), event1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, event2)
}
