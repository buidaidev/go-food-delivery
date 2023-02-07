package tokenprovider

import (
	"errors"
	"go-food-delivery/common"
	"time"
)

type Provider interface {
	Generate(data TokenPayload, expiry int) (*Token, error)
	Validate(token string) (*TokenPayload, error)
}

type Token struct {
	Token   string    `json:"token"`
	Created time.Time `json:"created"`
	Expiry  int       `json:"expiry"`
}

type TokenPayload struct {
	UserId int    `json:"user_id"`
	Role   string `json:"role"`
}

var (
	ErrTokenNotFound = common.NewCustomError(
		errors.New("token not found"),
		"token not found",
		"ErrorTokenNotFound",
	)
	ErrEncodingToken = common.NewCustomError(
		errors.New("error encoding token"),
		"error encoding token",
		"ErrorEncodingToken",
	)
	ErrInvalidToken = common.NewCustomError(
		errors.New("invalid token provider"),
		"invalid token provider",
		"ErrorInvalidToken",
	)
)
