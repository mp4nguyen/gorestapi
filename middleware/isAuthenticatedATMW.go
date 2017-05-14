package middleware

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"bitbucket.org/restapi/models/accessTokenMdl"
)

func IsAuthenticatedATMW(inner http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		accessTokenHeader := r.Header["Accesstoken"]
		//fmt.Println("IsAuthenticatedATMW: accessTokenHeader = ", accessTokenHeader)

		if len(accessTokenHeader) > 0 {
			accessToken := r.Header["Accesstoken"][0]
			//fmt.Println("IsAuthenticatedATMW: Accesstoken = ", accessToken)

			//get accessToken object in Redis
			at, err := accessTokenMdl.One(accessToken)
			//fmt.Println("AccessToken = ", at, err)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, "Unauthorized access to this resource(Token is invalid %s)", err)
			} else if at.Ttl.Before(time.Now().UTC()) {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprint(w, "Unauthorized access to this resource(Token is expired)")
			} else {
				//Checked authen successfully and then go to the main funcs
				//Add data to context here.
				ctx := context.WithValue(r.Context(), "UserId", at.UserId)
				ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
				defer cancel()
				inner.ServeHTTP(w, r.WithContext(ctx))
				//inner.ServeHTTP(w, r)
			}

		} else {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Unauthorized access to this resource; err = cannot find Accesstoken in header")
		}
	})
}
