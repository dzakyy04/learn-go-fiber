package handler

import (
	"fmt"
	"learn-go-fiber/database"
	"learn-go-fiber/model/entity"
	"learn-go-fiber/model/request"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func BookHandlerCreate(ctx *fiber.Ctx) error {
	// Get book request and parse body
	book := new(request.BookCreateRequest)
	if err := ctx.BodyParser(book); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse request body",
			"error":   err.Error(),
		})
	}

	// Validation
	var validate = validator.New()
	errValidate := validate.Struct(book)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Failed to validate request body",
			"error":   errValidate.Error(),
		})
	}

	// Handle require image
	filename := ctx.Locals("filename")
	if filename == nil {
		return ctx.Status(422).JSON(fiber.Map{
			"success": false,
			"message": "Failed to upload file",
			"error":   "Cover is required",
		})
	}

	filenameString := fmt.Sprintf("%v", filename)

	// Create book
	newBook := entity.Book{
		Title:  book.Title,
		Author: book.Author,
		Cover:  filenameString,
	}

	errCreateBook := database.DB.Create(&newBook).Error
	if errCreateBook != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to create book",
			"error":   errCreateBook.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "Book created successfully",
		"data":    newBook,
	})
}
