package jwtMiddleware

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

var jwtKey = []byte("SuperStrongPasswordMayBeNot!")

type Claims struct {
	UserId    string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	jwt.StandardClaims
}

func VerifyJSONWebToken(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	// fmt.Println(token)
	if token == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, bson.M{})
		return
	}
	tokenStrings := strings.Split(token, " ")
	if len(tokenStrings) != 2 {
		c.AbortWithStatusJSON(http.StatusBadRequest, bson.M{})
		return
	}
	tokenString := tokenStrings[1]

	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	// fmt.Println("token ", claims.UserId)
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, bson.M{})
			return
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, bson.M{})
		return
	}
	if !tkn.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, bson.M{})
		return
	}

	// if claims, ok := tkn.Claims.(jwt.MapClaims); ok && tkn.Valid {
	// 	fmt.Println("claims", claims)
	// }
	// fmt.Println(tkn.Claims.Valid())

	// fmt.Println("username ", claims.Id)
	c.Set("userId", claims.UserId)
	// fmt.Println("Successfully verified")
	c.Next()
}
