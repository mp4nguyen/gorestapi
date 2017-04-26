package myjwt

import (
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
)

func JWTMW(h http.HandlerFunc) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		accessToken := r.Header["Accesstoken"][0]
		fmt.Println("JWT Middleware: Accesstoken = ", accessToken)

		token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			// if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			//     return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			// }

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return VerifyKey, nil
		})

		if err == nil {
			if token.Valid {

				if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
					fmt.Println(claims["username"])
				} else {
					fmt.Println(err)
				}

				h.ServeHTTP(w, r)
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprint(w, "Token is not valid")
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Unauthorized access to this resource")
		}

	})
}
