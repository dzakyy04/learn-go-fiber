package utils

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func HandleSingleFile(ctx *fiber.Ctx) error {
	var filename *string

	file, errFile := ctx.FormFile("cover")
	if errFile != nil {
		if !strings.Contains(errFile.Error(), "no uploaded file") {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": "Failed to upload file",
				"error":   errFile.Error(),
			})
		}
	}

	if file != nil {
		filename = &file.Filename

		errSaveFile := ctx.SaveFile(file, fmt.Sprintf("./public/cover/%s", *filename))
		if errSaveFile != nil {
			return ctx.Status(500).JSON(fiber.Map{
				"success": false,
				"message": "Failed to save file",
				"error":   errSaveFile.Error(),
			})
		}
	}

	if filename != nil {
		ctx.Locals("filename", *filename)
	} else {
		ctx.Locals("filename", nil)
	}

	return ctx.Next()
}

func HandleMultipleFile(ctx *fiber.Ctx) error {
	form, errForm := ctx.MultipartForm()
	if errForm != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Failed to upload file",
			"error":   errForm.Error(),
		})
	}

	files := form.File["photos"]

	var filenames []string
	for i, file := range files {
		var filename string

		if file != nil {
			filename = fmt.Sprintf("%d-%v", i, file.Filename)

			errSaveFile := ctx.SaveFile(file, fmt.Sprintf("./public/photo/%s", filename))
			if errSaveFile != nil {
				return ctx.Status(500).JSON(fiber.Map{
					"success": false,
					"message": "Failed to save file",
					"error":   errSaveFile.Error(),
				})
			}
		}

		if filename != "" {
			filenames = append(filenames, filename)
		}
	}

	ctx.Locals("filenames", filenames)
	return ctx.Next()
}
