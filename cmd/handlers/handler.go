package handlers

import (
	"kissandeat/internal/service"
	"kissandeat/middleware"
	"log"

	"github.com/gin-gonic/gin"
	// swaggerFiles "github.com/swaggo/files"
	// ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	secretKey        string
	refreshSecretKey string
	services         *service.Service
}

func NewHandler(secretKey, refreshSecretKey string, services *service.Service) *Handler {
	return &Handler{secretKey: secretKey,
		refreshSecretKey: refreshSecretKey,
		services:         services}
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
		auth.POST("/login", func(c *gin.Context) {
			log.Println("Received /auth/login request")
			h.Login(c)
		})
		auth.POST("/register", h.Register)
		auth.POST("/refresh", h.RefreshToken)
		auth.POST("/family", h.AddFamily)
	}
	api := router.Group("/api")
	// User endpoints
	users := api.Group("/users")
	{
		// users.Use(middleware.JWTMiddleware("your_secret_key"))
		users.GET("", h.ListUsers)   // List users
		users.GET("/:id", h.GetUser) // Get user by ID
	}
	family := api.Group("/family")
	{
		// family.Use(middleware.JWTMiddleware("your_secret_key"))
		family.GET("", h.ListFamilies)  // List family
		family.GET("/:id", h.GetFamily) // family user by ID
	}
	// Dish endpoints
	dishes := api.Group("/dishes")
	{
		dishes.Use(middleware.JWTMiddleware("your_secret_key"))
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
