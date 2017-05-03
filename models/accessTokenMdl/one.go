package accessTokenMdl

import (
	"encoding/json"

	"bitbucket.org/restapi/db"
)

func One(accessToken string) (at AccessToken, err error) {

	atCaching := ATCaching{}

	atObj, err := atCaching.GetAccessToken(accessToken)

	//fmt.Println("atObj from caching = ", atObj, err)
	if err != nil {

		var deserialized AccessToken

		serialized, err := db.GetRedis().Get(accessToken).Result()

		//fmt.Println("serialized = ", serialized, " err = ", err)

		if err != nil {
			return deserialized, err
		} else {
			//fmt.Println("2serialized = ", serialized, " err = ", err)
			serializedInByte := []byte(serialized)

			err = json.Unmarshal(serializedInByte, &deserialized)

			//fmt.Println("serialized = ", deserialized, " err = ", err)

			if deserialized.AccessToken != "" {
				atCaching.SetAccessToken(deserialized.AccessToken, deserialized)
			}
			return deserialized, err
		}
	} else {
		return atObj, nil
	}
}
