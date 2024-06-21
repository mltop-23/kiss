package handlers

import (
	"context"
	"kissandeat/internal/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello, world!"})
}

func (h *Handler) Register(c *gin.Context) {
	var input structs.User

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.services.AuthInterface.CreateMember(c.Request.Context(), &input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "registration successful"})
}

func (h *Handler) Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.services.AuthInterface.LoginMember(c.Request.Context(), input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *Handler) ListUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "List users stub"})
}

func (h *Handler) GetUser(c *gin.Context) {

}

func (h *Handler) AddDish(c *gin.Context) {
	var dish structs.Dish
	if err := c.ShouldBindJSON(&dish); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := context.Background()
	err := h.services.DishInterface.AddDish(ctx, &dish)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add dish"})
		return
	}

	c.JSON(http.StatusOK, dish)
}

func (h *Handler) ListDishes(c *gin.Context) {
	ctx := context.Background()
	dishes, err := h.services.DishInterface.GetDishes(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get dishes"})
		return
	}

	c.JSON(http.StatusOK, dishes)
}

func (h *Handler) GetDish(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid dish ID"})
		return
	}

	ctx := context.Background()
	dish, err := h.services.DishInterface.GetDish(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get dish"})
		return
	}
	if dish == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Dish not found"})
		return
	}

	c.JSON(http.StatusOK, dish)
}
