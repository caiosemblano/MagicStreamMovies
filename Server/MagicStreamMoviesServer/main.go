package main

import (
	"fmt"

	//controller "github.com/caiosemblano/MagicStreamMovies/Server/MagicStreamMoviesServer/controllers"
	"github.com/gin-gonic/gin"
	"github.com/caiosemblano/MagicStreamMovies/Server/MagicStreamMoviesServer/routes"
)

func main() {
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello, Magic Stream Movies!")
	})
	routes.SetupProtectedRoutes(router)
	routes.SetupUnProtectedRoutes(router)


	if err := router.Run(":8080"); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
