package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/BigListRyRy/harbourlivingapi/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Phone:       sql.NullString{String: util.RandomName()},
		FirstName:   util.RandomName(),
		LastName:    util.RandomName(),
		Email:       util.RandomEmail(),
		Password:    util.RandomName(),
		Username:    util.RandomString(6),
		Usertype:    1,
		DateOfBirth: time.Date(1989, 8, 25, 0, 0, 0, 0, time.UTC),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Phone, user.Phone)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.FirstName, user.FirstName)
	require.Equal(t, arg.LastName, user.LastName)
	require.Equal(t, arg.Password, user.Password)
	require.Equal(t, arg.Username, user.Username)
	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.ID)
	require.NoError(t, err)
	require.Equal(t, user1.Phone, user2.Phone)
	require.Equal(t, user1.FirstName, user2.FirstName)
	require.Equal(t, user1.LastName, user2.LastName)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.Username, user2.Username)
}

func TestGetAllUser(t *testing.T) {

	for i := 0; i < 10; i++ {
		createRandomUser(t)
	}
	users, err := testQueries.GetAllUsers(context.Background())
	require.NoError(t, err)
	for _, user := range users {
		require.NotEmpty(t, user)
	}
}

func TestDeleteUser(t *testing.T) {
	user1 := createRandomUser(t)

	err := testQueries.DeleteUser(context.Background(), user1.ID)
	require.NoError(t, err)
	user2, err := testQueries.GetUser(context.Background(), user1.ID)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, user2)
}
