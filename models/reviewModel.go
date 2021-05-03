package models

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// type Review struct {
// 	id          string
// 	movieId     string
// 	description string
// 	rating      int
// 	userId      string
// }

func GetReview(movieId string) bson.M {
	movieObjectId, err := primitive.ObjectIDFromHex(movieId)

	if err != nil {
		log.Fatal(err)
	}

	tempContext := context.TODO()
	singleResult := reviewCollection.FindOne(tempContext, bson.M{"_id": movieObjectId})
	if singleResult.Err() != nil {
		log.Fatal(err)
	}
	var result bson.M

	if err = singleResult.Decode(&result); err != nil {
		log.Fatal(err)
	}
	// fmt.Println("result ", result)
	return result
}

func GetMovieReviews(movieId string) []bson.M {
	movieObjectId, err := primitive.ObjectIDFromHex(movieId)

	if err != nil {
		log.Fatal(err)
	}
	matchStage := bson.D{{"$match", bson.D{{"movieId", movieObjectId}}}}
	lookupStage := bson.D{{"$lookup", bson.D{{"from", "users"}, {"localField", "userId"}, {"foreignField", "_id"}, {"as", "reviewedBy"}}}}
	unwindStage := bson.D{{"$unwind", bson.D{{"path", "$reviewedBy"}, {"preserveNullAndEmptyArrays", false}}}}
	projectStage := bson.D{{"$project",
		bson.D{{"_id", 1},
			{"movieId", 1},
			{"description", 1},
			{"rating", 1},
			{"firstName", "$reviewedBy.firstName"},
			{"lastName", "$reviewedBy.lastName"},
		}}}

	tempContext := context.TODO()
	// var temp bson.M
	// cursor, err := reviewCollection.Find(tempContext, bson.M{"movieId": movieObjectId})
	cursor, err := reviewCollection.Aggregate(tempContext, mongo.Pipeline{matchStage, lookupStage, unwindStage, projectStage})

	var movieReviews []bson.M
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(tempContext)

	for cursor.Next(tempContext) {
		var result bson.M

		if err = cursor.Decode(&result); err != nil {
			log.Fatal(err)
		}
		// fmt.Println(result)
		movieReviews = append(movieReviews, result)
	}
	return movieReviews
}

func PrintAllReviews() {
	tempContext := context.TODO()
	cursor, err := reviewCollection.Find(tempContext, bson.M{})

	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(tempContext)

	for cursor.Next(tempContext) {
		var result bson.M

		if err = cursor.Decode(&result); err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)
	}
}

func CreateReview(movie string, user string, rating int, description string) error {
	userId, err := primitive.ObjectIDFromHex(user)
	if err != nil {
		return err
	}
	movieId, err := primitive.ObjectIDFromHex(movie)
	if err != nil {
		return err
	}
	tempContext := context.TODO()
	review, err := reviewCollection.InsertOne(tempContext, bson.D{
		{Key: "movieId", Value: movieId},
		{Key: "description", Value: description},
		{Key: "userId", Value: userId},
		{Key: "rating", Value: rating},
	})
	fmt.Println("created review ", review)
	return err
}
