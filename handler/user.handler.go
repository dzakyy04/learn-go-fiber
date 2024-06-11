package handler

import (
	"learn-go-fiber/database"
	"learn-go-fiber/model/entity"
	"log"

	"github.com/gofiber/fiber/v2"
)

func UserHandlerGetAll(ctx *fiber.Ctx) error {
	var users []entity.User

	result := database.DB.Find(&users)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return ctx.JSON(users)
}
