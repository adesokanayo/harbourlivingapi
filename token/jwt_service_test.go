package token

import (
	"testing"
	"time"

	"github.com/BigListRyRy/harbourlivingapi/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/require"
)

func TestJWTMaker(t *testing.T) {
	maker, err := NewJWTService(util.RandomString(32))
	require.NoError(t, err)

	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	userinfo := UserInfo{
		UserID:   1,
		Username: util.RandomOwner(),
		Email:    util.RandomEmail(),
		UserType: 1,
	}

	token, err := maker.CreateToken(userinfo, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	require.NotZero(t, payload.ID)
	require.Equal(t, userinfo.Username, payload.Username)
	require.Equal(t, userinfo.Email, payload.Email)
	require.Equal(t, userinfo.UserType, payload.UserType)
	require.Equal(t, userinfo.UserID, payload.UserID)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

func TestExpiredJWTToken(t *testing.T) {
	maker, err := NewJWTService(util.RandomString(32))
	require.NoError(t, err)

	userinfo := UserInfo{
		UserID:   1,
		Username: util.RandomOwner(),
		Email:    util.RandomEmail(),
		UserType: 1,
	}

	token, err := maker.CreateToken(userinfo, -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrExpiredToken.Error())
	require.Nil(t, payload)
}

func TestInvalidJWTTokenAlgNone(t *testing.T) {

	userinfo := UserInfo{
		UserID:   1,
		Username: util.RandomOwner(),
		Email:    util.RandomEmail(),
		UserType: 1,
	}
	payload, err := NewPayload(userinfo, time.Minute)
	require.NoError(t, err)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
	token, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err)

	maker, err := NewJWTService(util.RandomString(32))
	require.NoError(t, err)

	payload, err = maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrInvalidToken.Error())
	require.Nil(t, payload)
}

func TestParseToken(t *testing.T) {
	secret := util.RandomString(32)
	maker, err := NewJWTService(secret)
	require.NoError(t, err)

	userinfo := UserInfo{
		UserID:   1,
		Username: util.RandomOwner(),
		Email:    util.RandomEmail(),
		UserType: 1,
	}

	token, err := maker.CreateToken(userinfo, time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	username, err := maker.ParseToken(token)
	require.NoError(t, err)
	require.NotNil(t, username)
	require.Equal(t, userinfo.Username, username)
}
