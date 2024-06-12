package handler

import (
	"learn-go-fiber/database"
	"learn-go-fiber/model/entity"
	"learn-go-fiber/model/request"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func UserHandlerGetAll(ctx *fiber.Ctx) error {
	var users []entity.User

	result := database.DB.Find(&users)
	if result.Error != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Failed to retrieve users",
			"error":   result.Error.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"success": true,
		"message": "Users retrieved successfully",
		"data":    users,
	})
}

func UserHandlerGetById(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")

	var user entity.User
	err := database.DB.First(&user, userId).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.Status(404).JSON(fiber.Map{
				"status":  "error",
				"message": "User not found",
				"error":   err.Error(),
			})
		}
		return ctx.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to retrieve user",
			"error":   err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"status":  "success",
		"message": "User retrieved successfully",
		"data":    user,
	})
}
func UserHandlerCreate(ctx *fiber.Ctx) error {
	user := new(request.UserCreateRequest)
	if err := ctx.BodyParser(user); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse request body",
			"error":   err.Error(),
		})
	}

	var validate = validator.New()
	errValidate := validate.Struct(user)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Failed to validate request body",
			"error":   errValidate.Error(),
		})
	}

	newUser := entity.User{
		Name:    user.Name,
		Email:   user.Email,
		Address: user.Address,
		Phone:   user.Phone,
	}

	errCreateUser := database.DB.Create(&newUser).Error
	if errCreateUser != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Failed to create user",
			"error":   errCreateUser.Error(),
		})
	}

	return ctx.Status(201).JSON(fiber.Map{
		"success": true,
		"message": "User created successfully",
		"data":    newUser,
	})
}

func UserHandlerUpdate(ctx *fiber.Ctx) error {
	userRequest := new(request.UserUpdateRequest)
	if err := ctx.BodyParser(userRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse request body",
			"error":   err.Error(),
		})
	}

	userId := ctx.Params("id")
	var user entity.User

	err := database.DB.First(&user, userId).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.Status(404).JSON(fiber.Map{
				"status":  "error",
				"message": "User not found",
				"error":   err.Error(),
			})
		}
		return ctx.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to retrieve user",
			"error":   err.Error(),
		})
	}

	if userRequest.Name != "" {
		user.Name = userRequest.Name
	}
	if userRequest.Address != "" {
		user.Address = userRequest.Address
	}
	if userRequest.Phone != "" {
		user.Phone = userRequest.Phone
	}

	errUpdateUser := database.DB.Save(&user).Error
	if errUpdateUser != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Failed to update user",
			"error":   errUpdateUser.Error(),
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "User updated successfully",
		"data":    user,
	})
}