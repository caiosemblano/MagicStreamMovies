package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	controller "github.com/caiosemblano/MagicStreamMovies/Server/MagicStreamMoviesServer/controllers"
)

func main() {
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context){
		c.String(200, "Hello, Magic Stream Movies!")
	})

	router.GET("/movies", controller.GetMovies())
	router.GET("/movie/:imdb_id", controller.GetMovie())
	router.POST("/addmovie", controller.AddMovie())

	if err := router.Run(":8080"); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}