package middlewares

import (
	"os"
	"strings"

	"github.com/mesxx/Fiber_User_Management_API/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func RestrictedUser(c *fiber.Ctx) error {
	secret_key := os.Getenv("SECRET_KEY")

	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "invalid token")
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	token, err := jwt.ParseWithClaims(tokenString, &models.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret_key), nil
	})

	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}

	claims, ok := token.Claims.(*models.CustomClaims)
	if !ok {
		return fiber.NewError(fiber.StatusUnauthorized, "unknown claims type, cannot proceed")
	}

	c.Locals("user", claims)

	return c.Next()
}
