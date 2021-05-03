package movieRouter

import (
	"GoLang-WebServer/controller/movieController"

	"github.com/gin-gonic/gin"
)

// ListMoviesAPI /api/movies/ <GET>
// SearchMovieAPI /api/movies?q=<movieName> <POST>
// DisplayMovieAPI /api/movies/:id <GET>

func MovieRouter(rg *gin.RouterGroup) {
	movie := rg.Group("/")

	// Include this if the json web token is to be verified. Can be included to any route.
	// movie.GET("/", jwtMiddleware.VerifyJSONWebToken, movieController.Index)

	movie.GET("/", movieController.Index)

	movie.POST("/", movieController.SearchMovie)

	movie.GET("/:id", movieController.DisplayMovie)

}
