package handlers

import (
	"context"
	"errors"
	"kissandeat/internal/structs"
	"kissandeat/middleware"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
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
	log.Println("start login")
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := h.services.AuthInterface.LoginMember(c.Request.Context(), input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	id, err := strconv.Atoi(userID)
	accessToken, err := middleware.CreateToken(h.secretKey, id, time.Minute*1) // 15 минут для access токена
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create access token"})
		return
	}

	refreshToken, err := middleware.CreateRefreshToken(h.refreshSecretKey, id, time.Hour*24*7) // 7 дней для refresh токена
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create refresh token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func (h *Handler) RefreshToken(c *gin.Context) {
	var input struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	refreshToken, err := jwt.ParseWithClaims(input.RefreshToken, &structs.Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(h.refreshSecretKey), nil
	})

	if err != nil || !refreshToken.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid refresh token"})
		return
	}

	if claims, ok := refreshToken.Claims.(*structs.Claims); ok && refreshToken.Valid {
		if time.Now().Unix() > claims.ExpiresAt {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "refresh token expired"})
			return
		}

		newAccessToken, err := middleware.CreateToken(h.secretKey, claims.UserID, time.Minute*15) // 15 минут для access токена
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create new access token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"access_token": newAccessToken})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid refresh token claims"})
	}
}

// костыль
func (h *Handler) ListUsers(c *gin.Context) {
	ctx := context.Background()
	users, err := h.services.AuthInterface.GetFamily(ctx, 5)
	if err != nil {
		log.Printf("Error listing users: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list users"})
		return
	}

	c.JSON(http.StatusOK, users)
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

// func RefreshTokenHandler(secretKey string, refreshSecret string) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var request struct {
// 			RefreshToken string `json:"refresh_token" binding:"required"`
// 		}

// 		if err := c.ShouldBindJSON(&request); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		refreshToken, err := jwt.ParseWithClaims(request.RefreshToken, &structs.Claims{}, func(token *jwt.Token) (interface{}, error) {
// 			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 				return nil, errors.New("unexpected signing method")
// 			}
// 			return []byte(refreshSecret), nil
// 		})

// 		if err != nil || !refreshToken.Valid {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid refresh token"})
// 			return
// 		}

// 		if claims, ok := refreshToken.Claims.(*structs.Claims); ok && refreshToken.Valid {
// 			if time.Now().Unix() > claims.ExpiresAt {
// 				c.JSON(http.StatusUnauthorized, gin.H{"error": "refresh token expired"})
// 				return
// 			}

// 			// Создаем новый access токен
// 			newAccessToken, err := middleware.CreateToken(secretKey, claims.UserID, time.Minute*15) // 15 минут для access токена
// 			if err != nil {
// 				c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create new access token"})
// 				return
// 			}

// 			c.JSON(http.StatusOK, gin.H{"access_token": newAccessToken})
// 		} else {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid refresh token claims"})
// 		}
// 	}
// }
