package movieController

import (
	"GoLang-WebServer/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ListMoviesAPI /api/movies/ <GET>
// SearchMovieAPI /api/movies?q=<movieName> <POST>
// DisplayMovieAPI /api/movies/:id <GET>

type SearchMovieFormat struct {
	Name string `form:"movieName" json:"movieName" binding:"required"`
}

func Index(c *gin.Context) {
	movieList := models.FindAll()
	// fmt.Print(movieList)
	c.JSON(http.StatusOK, movieList)
}

func DisplayMovie(c *gin.Context) {
	//fmt.Println(c.Params)
	movieId := c.Param("id")
	fmt.Println(movieId)
	if movieId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Provide ID"})
		return
	}
	movieObject := models.ShowMovie(movieId)
	fmt.Println(movieObject)
	// c.JSON(http.StatusOK, gin.H{})
	c.JSON(http.StatusOK, movieObject)
}

func SearchMovie(c *gin.Context) {

	var movieData SearchMovieFormat
	c.BindJSON(&movieData)

	movieName := movieData.Name

	if movieName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please fill the field: movie name"})
		return
	}
	movieObject := models.FindMovie(movieName)
	// fmt.Println(len(movieObject))
	if len(movieObject) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No matches found"})
	} else {
		c.JSON(http.StatusOK, movieObject)
	}

}
