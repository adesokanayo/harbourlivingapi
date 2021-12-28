package db

import (
	"context"
	"database/sql"
	"log"
	"testing"

	"github.com/BigListRyRy/harbourlivingapi/util"
	"github.com/stretchr/testify/require"
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
		Status:      1,
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
	require.NotZero(t, event.ID)
	require.NotZero(t, event.StartDate)
	require.NotZero(t, event.EndDate)
	require.NotZero(t, event.CreatedAt)
	require.Equal(t, arg.Status, event.Status)

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
}

func TestUpdateEvent(t *testing.T) {
	event1 := createRandomEvent(t)

	startdate, err := util.ProcessDateTime("rfc", "2015-09-15T14:00:12-00:00")
	if err != nil {
		log.Fatal(err)
	}
	enddate, err := util.ProcessDateTime("rfc", "2015-09-15T14:00:12-00:00")
	arg := UpdateEventParams{
		TitleToUpdate:       true,
		Title:               "Updated Title",
		DescriptionToUpdate: true,
		Description:         "new description",
		BannerImage:         "new banner image",
		BannerImageToUpdate: true,
		StartDate:           *startdate,
		StartDateToUpdate:   true,
		EndDate:             *enddate,
		EndDateToUpdate:     true,
		VenueToUpdate:       true,
		Venue:               1,
		CategoryToUpdate:    true,
		Category:            1,
		TypeToUpdate:        true,
		Type:                1,
		StatusToUpdate:      true,
		Status:              1,
		ID:                  event1.ID,
	}
	_, err = testQueries.UpdateEvent(context.Background(), arg)
	updatedEvent1, err := testQueries.GetEvent(context.Background(), event1.ID)
	require.NoError(t, err)
	require.Equal(t, arg.Title, updatedEvent1.Title)
	require.Equal(t, arg.Description, updatedEvent1.Description)
	require.Equal(t, arg.BannerImage, updatedEvent1.BannerImage)

	//require.Equal(t, arg.StartDate, updatedEvent1.StartDate)
	//require.Equal(t, arg.EndDate, updatedEvent1.EndDate)
	require.Equal(t, arg.Venue, updatedEvent1.Venue)
	require.Equal(t, arg.Type, updatedEvent1.Type)
	require.Equal(t, arg.Status, updatedEvent1.Status)

}

func TestGetAllEvents(t *testing.T) {

	for i := 0; i < 10; i++ {
		createRandomEvent(t)
	}

	arg := GetEventsParams{
		Limit:       1,
		Offset:      1,
		Category:    1,
	}
	events, err := testQueries.GetEvents(context.Background(), arg)
	require.NoError(t, err)
	for _, event := range events {
		require.NotEmpty(t, event)
	}
}

func TestDeleteEvent(t *testing.T) {
	event1 := createRandomEvent(t)
	err := testQueries.DeleteEvent(context.Background(), event1.ID)
	require.NoError(t, err)
	event2, err := testQueries.GetEvent(context.Background(), event1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, event2)
}
