package accessTokenMdl

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"io"
	"time"

	"bitbucket.org/restapi/db"
)

func Create(userId int) (accessToken string, err error) {

	token := make([]byte, TokenLength)
	if _, err := io.ReadFull(rand.Reader, token); err != nil {
		return "", err
	}

	at := AccessToken{
		UserId:      userId,
		AccessToken: base64.URLEncoding.EncodeToString(token),
		Ttl:         time.Now().UTC().Add(TtlDuration),
		CreatedAt:   time.Now().UTC(),
	}

	serialized, err := json.Marshal(at)
	if err != nil {
		return "", err
	}
	_, errRedis := db.GetRedis().Set(at.AccessToken, serialized, 0).Result()
	if errRedis != nil {
		return "", errRedis
	}

	return at.AccessToken, err
}
