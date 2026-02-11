package routes

import (
	controller "github.com/caiosemblano/MagicStreamMovies/Server/MagicStreamMoviesServer/controllers"
	"github.com/gin-gonic/gin"
	"github.com/caiosemblano/MagicStreamMovies/Server/MagicStreamMoviesServer/middleware"
)

func SetupProtectedRoutes(router *gin.Engine) {
	router.Use(middleware.AuthMiddleware())
	router.POST("/addmovie", controller.AddMovie())
	router.GET("/movie/:imdb_id", controller.GetMovie())
	
}