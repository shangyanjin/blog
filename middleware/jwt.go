package middleware

import (
	"blog/config"
	"blog/model"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var (
	// Change this secret in production
	secretKey = []byte(config.GetString("jwt.secret", "your-secret-key"))

	// ErrMissingAuthHeader is returned when auth header is missing
	ErrMissingAuthHeader = errors.New("missing authorization header")
	// ErrInvalidAuthHeader is returned when auth header is invalid
	ErrInvalidAuthHeader = errors.New("invalid authorization header")
)

// Claims represents the JWT claims
type Claims struct {
	UserID int    `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// GenerateToken generates a new JWT token
func GenerateToken(userID int, role string) (string, error) {
	claims := Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // Token expires in 24 hours
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

// JWTAuth middleware for JWT authentication
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": ErrMissingAuthHeader.Error(),
			})
			return
		}

		// Bearer token format
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": ErrInvalidAuthHeader.Error(),
			})
			return
		}

		tokenString := parts[1]
		claims := &Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}

		if !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token",
			})
			return
		}

		// Set claims to context
		c.Set("userID", claims.UserID)
		c.Set("role", claims.Role)

		c.Next()
	}
}

// GetUserID gets the user ID from the context
func GetUserID(c *gin.Context) uint {
	userID, exists := c.Get("userID")
	if !exists {
		return 0
	}
	return userID.(uint)
}

// GetUserRole gets the user role from the context
func GetUserRole(c *gin.Context) string {
	role, exists := c.Get("role")
	if !exists {
		return ""
	}
	return role.(string)
}

// GenerateJWT generates a JWT token
func GenerateJWT(userId int, userName string, userRole string) (string, error) {
	const tokenExpiration = 24 * time.Hour

	if len(userRole) < 1 {
		return "", errors.New("userType cannot be empty")
	}

	var user model.User
	if err := model.DB.Model(&model.User{}).Where("id = ?", userId).First(&user).Error; err != nil {
		logrus.Errorf("Failed to GenerateJWT user detail for ID %d: %v", userId, err)
		return "", err
	}

	claims := Claims{
		UserID: userId,
		Role:   userRole,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // Token expires in 24 hours
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.GetString("server.jwt_secret_key", "mix-jwt-secret-key")))
}

// parseJWT parses and validates the JWT token
func parseJWT(tokenString string) (*Claims, error) {
	if len(tokenString) < 32 {
		return nil, errors.New("token length is too short")
	}

	// Get JWT secret key from config
	secretKey := []byte(config.GetString("server.jwt_secret_key", "mix-jwt-secret-key"))

	// Parse token with custom validation
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Validate signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %v", err)
	}

	// Type assert and validate claims
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token claims")
	}

	// Additional validation
	if claims.UserID <= 0 {
		return nil, errors.New("invalid user ID in token")
	}

	if claims.Role == "" {
		return nil, errors.New("missing role in token")
	}

	return claims, nil
}
