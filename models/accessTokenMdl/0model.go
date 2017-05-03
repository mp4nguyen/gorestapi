package accessTokenMdl

import (
	"time"
)

const (
	TokenLength int           = 64
	TtlDuration time.Duration = 20 * time.Minute
)

type AccessToken struct {
	UserId      int       `json:"userId"`
	AccessToken string    `json:"accessToken"`
	CreatedAt   time.Time `json:"createdAt"`
	Ttl         time.Time `json:"lifeTime"`
}

type AccessTokens struct {
	AccessTokens []AccessToken `json:"accessTokens"`
}

var accessTokenCaching map[string]AccessToken

type ATCaching struct{}
