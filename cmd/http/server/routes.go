package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initRoutes(router *gin.Engine, dependency *Dependency) {
	router.GET("/health", func(request *gin.Context) {
		request.String(http.StatusOK, "pong")
	})

	router.POST("/users", dependency.UserHandler.Create)
	router.GET("/users/:user_id", dependency.UserHandler.GetByID)
	router.GET("/users/", dependency.UserHandler.GetAll)
	router.PUT("/users/:user_id", dependency.UserHandler.Update)
	router.DELETE("/users/:user_id", dependency.UserHandler.Delete)
}
