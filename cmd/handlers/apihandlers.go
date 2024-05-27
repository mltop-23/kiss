package handlers

import (
	"encoding/json"
	"fmt"
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

// func GetUsersHandler(userService dbservice.Db) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		users, err := userService.GetUser(r.Context(),)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}

// 		// Преобразовать пользователей в JSON и отправить ответ
// 		jsonData, err := json.Marshal(users)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}

// 		w.Header().Set("Content-Type", "application/json")
// 		w.Write(jsonData)
// 	}
// }

func (h *Handler) GetUser(c *gin.Context) {

	// Получение ID пользователя из параметра
	// userID := c.Param("userID")

	// // Проверка валидности ID
	// fmt.Println(userID)
	// if _, err := uuid.Parse(userID); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
	// 	return
	// }
	// id, err := strconv.ParseInt(userID, 10, 64) // Parse userID as base-10 int64
	// if err != nil {
	// 	// Handle error gracefully (e.g., return bad request or validation error)
	// 	return
	// }
	id := int64(5)

	// Получение пользователя из сервиса
	// fmt.Println(id)
	user, err := h.services.GetUser(id)
	if err != nil {
		switch err.(type) {
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	fmt.Println(user)
	// fmt.Println(user)
	// Преобразование пользователя в JSON и отправка ответа
	jsonData, err := json.Marshal(user)
	// base64EncodedData := base64.StdEncoding.EncodeToString(jsonData)
	// fmt.Println(base64EncodedData)
	// fmt.Println("----------------------------------")
	// fmt.Println(jsonData)
	// err = json.Unmarshal(jsonData, &user)
	// fmt.Println("----------------------------------")
	// fmt.Println(jsonData)
	// fmt.Println(jsonData)
	// var userr structs.User
	err = json.Unmarshal(jsonData, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, jsonData)
}

func (h *Handler) ListDishes(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "List dishes stub"})
}

func (h *Handler) GetDish(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get dishes stub"})
}
