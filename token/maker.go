package token

import (
	"time"
)

type TokenService interface {
	CreateToken(userinfo UserInfo, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
	ParseToken(token string) (string, error)
}
