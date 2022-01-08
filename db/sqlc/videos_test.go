package db

import (
	"context"
	"database/sql"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_createVideo(t *testing.T) {
	//user := createRandomUser(t)
	arg := CreateVideoParams{

		Name: sql.NullString{String:"video name",Valid: true},
		Url:  "http://video.com",
	}
	video, err := testQueries.CreateVideo(context.Background(), arg)

	if err != nil {
		log.Fatal(err)
	}
	assert.Nil(t, err)
	assert.Equal(t, arg.Name, video.Name)
	assert.Equal(t, arg.Url , video.Url)

}

