package middleware

import (
	"learn-go-fiber/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(ctx *fiber.Ctx) error {
	// Check token from header
	authHeader := ctx.Get("Authorization")
	if authHeader == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "Unauthorized",
		})
	}

	// Split header to get token
	tokenString := strings.Split(authHeader, "Bearer ")[1]
	if tokenString == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "Unauthorized",
		})
	}

	// Parse and validate token
	token, err := utils.ParseToken(tokenString)
	if err != nil || !token.Valid {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "Invalid or expired token",
		})
	}

	// Set user information
	claims := token.Claims.(jwt.MapClaims)
	ctx.Locals("userID", claims["user_id"])

	return ctx.Next()
}
