package models

import (
	"context"
	"errors"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

type UserDetailFormat struct {
	UserId       string `bson:"_id"`
	FirstName    string `bson:"firstName"`
	LastName     string `bson:"lastName"`
	UserEmail    string `bson:"email"`
	UserPassword string `bson:"password"`
}

func FindAllUsers() []*UserDetailFormat {
	tempContext := context.TODO()
	cursor, err := userCollection.Find(tempContext, bson.M{})

	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(tempContext)

	var userList []*UserDetailFormat
	// fmt.Println("Trying to print all users ")
	for cursor.Next(tempContext) {
		var result UserDetailFormat

		if err = cursor.Decode(&result); err != nil {
			log.Fatal(err)
		}

		// To convert objectId to hex
		// user.UserId = result["u_id"].(primitive.ObjectID).Hex()
		// user.userName = result["u_name"].(string)
		// user.userEmail = result["u_email"].(string)
		// user.userPassword = result["u_password"].(string)

		userList = append(userList, &result)

	}
	return userList
}

func CreateUser(email string, password string, firstName string, lastName string) error {
	tempContext := context.TODO()
	_, err := userCollection.InsertOne(tempContext, bson.D{
		{Key: "email", Value: email},
		{Key: "firstName", Value: firstName},
		{Key: "lastName", Value: lastName},
		{Key: "password", Value: password},
	})
	// fmt.Println(result)
	return err
}

func VerifyLogin(email string, password []byte) (UserDetailFormat, error) {
	tempContext := context.TODO()
	result := userCollection.FindOne(tempContext, bson.D{
		{Key: "email", Value: email},
	})
	var isPresent UserDetailFormat
	err := result.Decode(&isPresent)
	if err != nil {
		return UserDetailFormat{}, errors.New("Invalid email or password")
	}
	fmt.Println(isPresent.UserPassword)
	err = bcrypt.CompareHashAndPassword([]byte(isPresent.UserPassword), password)
	if err != nil {
		return UserDetailFormat{}, errors.New("Invalid email or password")
	}
	return isPresent, nil
}
