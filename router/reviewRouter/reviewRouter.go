package reviewRouter

import (
	"GoLang-WebServer/controller/reviewController"
	"GoLang-WebServer/middlewares/jwtMiddleware"

	"github.com/gin-gonic/gin"
)

// ListReviewsAPI /api/review/ <GET>
// CreateReviewAPI /api/review/:id?r= <POST>
// UpdateReviewAPI /api/review/:rid?r= <PUT>
// DeleteReviewAPI /api/review/:rid <DELETE>

func ReviewRouter(rg *gin.RouterGroup) {
	review := rg.Group("/")
	// id - movie id
	// rid - review id

	review.GET("/:id", reviewController.ListReview)

	review.POST("/:id/", jwtMiddleware.VerifyJSONWebToken, reviewController.CreateReview)

	review.PUT("/:rid", reviewController.UpdateReview)

	review.DELETE("/:rid", reviewController.DeleteReview)

}
