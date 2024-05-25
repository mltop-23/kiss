package handlers

import (
	"kissandeat/internal/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	hello := router.Group("/hello")
	{
		hello.GET("", h.Hello) // Stub for hello endpoint
	}

	auth := router.Group("/auth")
	{
		auth.POST("/login", h.Login)       // Stub for login endpoint
		auth.POST("/register", h.Register) // Stub for register endpoint
	}

	users := router.Group("/users")
	{
		users.GET("", h.ListUsers)   // Stub for listing users
		users.GET("/:id", h.GetUser) // Stub for getting a user
	}

	dishes := router.Group("/dishes")
	{
		dishes.GET("", h.ListDishes)  // Stub for listing dishes
		dishes.GET("/:id", h.GetDish) // Stub for getting a dish
	}

	return router
}
