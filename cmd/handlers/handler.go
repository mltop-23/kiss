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

	// Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// router.GET("/", h.Index)

	api := router.Group("/api/v1")

	// Auth endpoints
	auth := api.Group("/auth")
	{
		auth.POST("/login", h.Login)       // Login endpoint
		auth.POST("/register", h.Register) // Register endpoint
	}

	// User endpoints
	users := api.Group("/users")
	{
		users.GET("", h.ListUsers)   // List users
		users.GET("/:id", h.GetUser) // Get user by ID
	}

	// Dish endpoints
	dishes := api.Group("/dishes")
	{
		dishes.GET("", h.ListDishes)  // List dishes
		dishes.GET("/:id", h.GetDish) // Get dish by ID
		// dishes.PUT("/:id", h.UpdateDish)   // Update dish by ID
		// dishes.DELETE("/:id", h.DeleteDish) // Delete dish by ID
	}

	// Apply middleware for authentication to endpoints that require it
	// users.Use(h.AuthMiddleware())
	// dishes.Use(h.AuthMiddleware())

	return router
}
