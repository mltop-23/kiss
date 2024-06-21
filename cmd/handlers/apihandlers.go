package handlers

import (
	"context"
	"kissandeat/internal/structs"
	"log"
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
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("Invalid user ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	ctx := context.Background()
	user, err := h.services.AuthInterface.GetMember(ctx, id)
	if err != nil {
		log.Printf("Error getting user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		return
	}
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) AddFamily(c *gin.Context) {
	var family structs.Family
	if err := c.ShouldBindJSON(&family); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := context.Background()

	// Проверка существования пользователей
	husbandExists, err := h.userExists(ctx, family.HusbandID)
	if err != nil {
		log.Printf("Error checking husband existence: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check husband existence"})
		return
	}
	if !husbandExists {
		log.Printf("Husband with ID %d does not exist", family.HusbandID)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Husband does not exist"})
		return
	}

	wifeExists, err := h.userExists(ctx, family.WifeID)
	if err != nil {
		log.Printf("Error checking wife existence: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check wife existence"})
		return
	}
	if !wifeExists {
		log.Printf("Wife with ID %d does not exist", family.WifeID)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wife does not exist"})
		return
	}

	err = h.services.AuthInterface.RegisterFamily(ctx, &family)
	if err != nil {
		log.Printf("Error registering family: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add family"})
		return
	}

	c.JSON(http.StatusOK, family)
}

func (h *Handler) userExists(ctx context.Context, userID int) (bool, error) {
	user, err := h.services.AuthInterface.GetMember(ctx, userID)
	if err != nil {
		return false, err
	}
	return user != nil, nil
}

func (h *Handler) ListFamilies(c *gin.Context) {
	ctx := context.Background()
	families, err := h.services.AuthInterface.ListFamilies(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list families"})
		return
	}

	c.JSON(http.StatusOK, families)
}

func (h *Handler) GetFamily(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid family ID"})
		return
	}

	ctx := context.Background()
	family, err := h.services.AuthInterface.GetFamily(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get family"})
		return
	}
	if family == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Family not found"})
		return
	}

	c.JSON(http.StatusOK, family)
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
