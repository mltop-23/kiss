package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello, world!"})
}

func (h *Handler) Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Login stub"})
}

func (h *Handler) Register(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Register stub"})
}

func (h *Handler) ListUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "List users stub"})
}

func (h *Handler) GetUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get user stub"})
}

func (h *Handler) ListDishes(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "List dishes stub"})
}

func (h *Handler) GetDish(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get dishes stub"})
}
