package middleware

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"kissandeat/internal/structs"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateToken(secretKey string, userID int, duration time.Duration) (string, error) {
	log.Println("start create")
	expirationTime := time.Now().Add(duration).Unix()
	claims := &structs.Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	log.Printf("Createdd Token: %s", signedToken) // Log the created token
	return signedToken, nil
}

func JWTMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("Entering JWT Middleware")
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header required"})
			c.Abort()
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header format must be Bearer {token}"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Log the token string for debugging
		log.Printf("Token String: %s", tokenString)

		token, err := jwt.ParseWithClaims(tokenString, &structs.Claims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secretKey), nil
		})

		if err != nil {
			// Log the error for debugging
			log.Printf("Error parsing token: %v", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		// Log token headers and claims
		log.Printf("Token Headers: %v", token.Header)
		log.Printf("Token Claims: %v", token.Claims)

		if claims, ok := token.Claims.(*structs.Claims); ok && token.Valid {
			if time.Now().Unix() > claims.ExpiresAt {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "token expired"})
				c.Abort()
				return
			}
			c.Set("userID", claims.UserID)
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token claims"})
			c.Abort()
			return
		}

		c.Next()
	}
}

// createRefreshToken creates a refresh token with the specified duration
func CreateRefreshToken(secretKey string, userID int, duration time.Duration) (string, error) {
	expirationTime := time.Now().Add(duration).Unix()
	claims := &structs.Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	log.Printf("refresh Token: %s", token)
	return token.SignedString([]byte(secretKey))
}
