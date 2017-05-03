package accessTokenMdl

import "errors"

func InitAccessTokenCaching() {
	accessTokenCaching = make(map[string]AccessToken)
}

func (c ATCaching) SetAccessToken(at string, accessToken AccessToken) {
	//fmt.Println("accessTokenCaching = ", accessTokenCaching, len(accessTokenCaching), accessTokenCaching == nil)

	if accessTokenCaching == nil {
		InitAccessTokenCaching()
	}
	_, ok := accessTokenCaching[at]
	if !ok {
		accessTokenCaching[at] = accessToken
	}
}

func (c ATCaching) GetAccessToken(at string) (accessToken AccessToken, err error) {
	atObject, ok := accessTokenCaching[at]
	if ok {
		return atObject, nil
	} else {
		return atObject, errors.New("AT has not existed in caching")
	}

}

func (c ATCaching) DeleteAccessToken(at string) {
	delete(accessTokenCaching, at)
}
