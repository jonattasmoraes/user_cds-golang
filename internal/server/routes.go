package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jonattasmoraes/app-go/internal/app/user/handler"
)

func initializeRoutes(router *gin.Engine) {
	// Initialize handlers
	handler.InitializeHandler()

	// Routes
	v1 := router.Group("/api")
	{
		v1.POST("/users", handler.CreateUserHandler)
		v1.GET("/users", handler.GetAllUsersHandler)
		v1.DELETE("/users/:id", handler.DeleteUserByIdHandler)
	}
}
