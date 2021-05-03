package apiRouter

import (
	"GoLang-WebServer/router/authRouter"
	"GoLang-WebServer/router/movieRouter"
	"GoLang-WebServer/router/reviewRouter"

	"github.com/gin-gonic/gin"
)

func RouteApi(rg *gin.RouterGroup) {

	auth := rg.Group("/auth")
	authRouter.AuthRouter(auth)

	movie := rg.Group("/movies")
	movieRouter.MovieRouter(movie)

	review := rg.Group("/review")
	reviewRouter.ReviewRouter(review)

}
