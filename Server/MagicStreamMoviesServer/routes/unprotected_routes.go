package routes

import (
	controller "github.com/caiosemblano/MagicStreamMovies/Server/MagicStreamMoviesServer/controllers"
	"github.com/gin-gonic/gin"

)

func SetupUnProtectedRoutes(router *gin.Engine) {
	router.POST("/register", controller.RegisterUser())
	router.POST("/login", controller.LoginUser())
	router.GET("/movies", controller.GetMovies())

}