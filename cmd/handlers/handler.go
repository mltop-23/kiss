package handlers

import (
	"kissandeat/internal/service"
	"kissandeat/middleware"

	"github.com/gin-gonic/gin"
	// swaggerFiles "github.com/swaggo/files"
	// ginSwagger "github.com/swaggo/gin-swagger"
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
	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// router.Use(middleware.JWTMiddleware("your-secret-key"))

	// router.GET("/", h.Index)

	// Auth endpoints
	auth := router.Group("/auth")
	{
		auth.POST("/login", h.Login)       // Login endpoint
		auth.POST("/register", h.Register) // Register endpoint
	}
	api := router.Group("/api")
	// User endpoints
	users := api.Group("/users")
	{
		users.Use(middleware.JWTMiddleware("your-secret-key"))
		users.GET("", h.ListUsers)   // List users
		users.GET("/:id", h.GetUser) // Get user by ID
	}

	// Dish endpoints
	dishes := api.Group("/dishes")
	{
		dishes.Use(middleware.JWTMiddleware("your-secret-key"))
		dishes.GET("", h.ListDishes)   // List dishes
		dishes.GET("/:id", h.GetDish)  // Get dish by ID
		dishes.POST("/:id", h.AddDish) // add dish by ID
		// dishes.DELETE("/:id", h.DeleteDish) // Delete dish by ID
	}

	// Apply middleware for authentication to endpoints that require it
	// users.Use(h.AuthMiddleware())
	// dishes.Use(h.AuthMiddleware())

	return router
}
