package server

import (
	"github.com/gin-gonic/gin"
	docs "github.com/jonattasmoraes/app-go/docs"
	"github.com/jonattasmoraes/app-go/internal/app/user/handler"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	// gin-swagger middleware
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
		v1.GET("/users/:id", handler.GetUserByIdHandler)
		v1.PUT("/users/:id", handler.UpdateUserByIdHandler)
		v1.PATCH("/users/:id", handler.PatchUserHandler)
	}

	// Swagger
	docs.SwaggerInfo.BasePath = "/api"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
