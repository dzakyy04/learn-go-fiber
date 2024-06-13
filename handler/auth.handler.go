package handler

import (
	"learn-go-fiber/database"
	"learn-go-fiber/model/entity"
	"learn-go-fiber/model/request"
	"learn-go-fiber/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func LoginHandler(ctx *fiber.Ctx) error {
	// Parse request body
	loginRequest := new(request.LoginRequest)
	if err := ctx.BodyParser(loginRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse request body",
			"error":   err.Error(),
		})
	}

	// Validate request body
	validate := validator.New()
	errValidate := validate.Struct(loginRequest)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Failed to validate request body",
			"error":   errValidate.Error(),
		})
	}

	// Check if user exists
	var user entity.User
	err := database.DB.Where("email = ?", loginRequest.Email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.Status(404).JSON(fiber.Map{
				"success": false,
				"message": "User not found",
				"error":   err.Error(),
			})
		}
		return ctx.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Failed to retrieve user",
			"error":   err.Error(),
		})
	}

	// Check password
	isValid := utils.CheckPasswordHash(loginRequest.Password, user.Password)

	if !isValid {
		return ctx.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Wrong credential",
		})
	}

	// Return token
	return ctx.JSON(fiber.Map{
		"success": true,
		"message": "User logged in successfully",
		"token":   "token",
		"data":    user,
	})
}
