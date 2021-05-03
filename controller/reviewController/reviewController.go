package reviewController

import (
	"GoLang-WebServer/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateReviewAPI /api/review/:id?r= <POST>
// UpdateReviewAPI /api/review/:rid?r= <PUT>
// ListReviewsAPI /api/review/ <GET>
// DeleteReviewAPI /api/review/:rid <DELETE>

type createReviewData struct {
	Description string `form:"description" json:"description" binding:"required"`
	Rating      int    `form:"rating" json:"rating" binding:"required"`
}

func CreateReview(c *gin.Context) {
	// err := jwtMiddleware.VerifyJWTAsFunc(c)
	// if err != nil {
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Unable to register review"})
	// }
	var review createReviewData
	c.BindJSON(&review)

	movieId := c.Param("id")
	UserId := c.MustGet("userId").(string)
	// fmt.Println("Got id", )
	// fmt.Println("movie id",)

	if movieId == "" || UserId == "" || review.Description == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Please provide both movie rating and description of review"})
	} else {
		err := models.CreateReview(movieId, UserId, review.Rating, review.Description)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Unable to register review"})
		}
		c.JSON(http.StatusOK, gin.H{})
	}

}

func UpdateReview(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func ListReview(c *gin.Context) {
	movieId := c.Param("id")
	// fmt.Println(movieId)
	result := models.GetMovieReviews(movieId)
	// models.PrintAllReviews()
	fmt.Println(result)
	c.JSON(http.StatusOK, result)
}

func DeleteReview(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
