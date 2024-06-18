package handler

import (
	"learn-go-fiber/database"
	"learn-go-fiber/model/entity"
	"learn-go-fiber/model/request"
	"learn-go-fiber/utils"

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
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"success": false,
			"message": "Failed to upload file",
			"error":   "Cover is required",
		})
	}

	filenamesData := filenames.([]string)
	for _, filename := range filenamesData {
		newPhoto := entity.Photo{
			Image:      filename,
			CategoryID: photo.CategoryID,
		}

		errCreatePhoto := database.DB.Create(&newPhoto).Error
		if errCreatePhoto != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"message": "Some data not saved properly",
				"error":   errCreatePhoto.Error(),
			})
		}
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "Photo created successfully",
	})
}

func PhotoHandlerDelete(ctx *fiber.Ctx) error {
	photoId := ctx.Params("id")
	var photo entity.Photo

	// Find photo
	errDelete := database.DB.First(&photo, photoId).Error
	if errDelete != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Photo not found",
			"error":   errDelete.Error(),
		})
	}

	// Delete file
	errDeleteFile := utils.HandleRemoveFile(photo.Image, "./public/photo/")
	if errDeleteFile != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to delete file",
			"error":   errDeleteFile.Error(),
		})
	}

	errDeletePhoto := database.DB.Delete(&photo).Error
	if errDeletePhoto != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to delete photo",
			"error":   errDeletePhoto.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Photo deleted successfully",
	})
}
