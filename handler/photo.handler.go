package handler

import (
	"fmt"
	"learn-go-fiber/model/request"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func PhotoHandlerCreate(ctx *fiber.Ctx) error {
	// Get book request and parse body
	photo := new(request.PhotoCreateRequest)
	if err := ctx.BodyParser(photo); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse request body",
			"error":   err.Error(),
		})
	}

	// Validation
	var validate = validator.New()
	errValidate := validate.Struct(photo)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Failed to validate request body",
			"error":   errValidate.Error(),
		})
	}

	// Handle require image
	filenames := ctx.Locals("filenames")
	if filenames == nil {
		return ctx.Status(422).JSON(fiber.Map{
			"success": false,
			"message": "Failed to upload file",
			"error":   "Cover is required",
		})
	}

	filenamesString := fmt.Sprintf("%v", filenames)
	
	log.Println(filenamesString)

	// Create book
	// newPhoto := entity.Photo{
	// 	Image:  filename,
	// 	CategoryID: 1,
	// }

	// errCreatePhoto := database.DB.Create(&newPhoto).Error
	// if errCreatePhoto != nil {
	// 	return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"success": false,
	// 		"message": "Failed to create photo",
	// 		"error":   errCreatePhoto.Error(),
	// 	})
	// }

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "Photo created successfully",
		// "data":    newBook,
	})
}
