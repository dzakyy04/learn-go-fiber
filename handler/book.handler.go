package handler

import (
	"fmt"
	"learn-go-fiber/database"
	"learn-go-fiber/model/entity"
	"learn-go-fiber/model/request"
	"strings"

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

	// File handler
	var filename string

	file, errFile := ctx.FormFile("cover")
	if errFile != nil {
		if strings.Contains(errFile.Error(), "no uploaded file") {
			filename = ""
		} else {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": "Failed to upload file",
				"error":   errFile.Error(),
			})
		}
	}

	if file != nil {
		filename = file.Filename

		errSaveFile := ctx.SaveFile(file, fmt.Sprintf("./public/cover/%s", filename))
		if errSaveFile != nil {
			return ctx.Status(500).JSON(fiber.Map{
				"success": false,
				"message": "Failed to save file",
				"error":   errSaveFile.Error(),
			})
		}
	}

	// Create book
	newBook := entity.Book{
		Title:  book.Title,
		Author: book.Author,
		Cover:  filename,
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
