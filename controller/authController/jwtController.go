package authController

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("SuperStrongPasswordMayBeNot!")

type Claims struct {
	UserId    string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	jwt.StandardClaims
}

func SignJWT(id string, firstName string, lastName string) (string, int) {
	// 5 minutes
	expirationTime := time.Now().Add(60 * time.Minute)

	claims := &Claims{
		UserId:    id,
		FirstName: firstName,
		LastName:  lastName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println("Token string", tokenString)
	return tokenString, int(expirationTime.Unix())
}
