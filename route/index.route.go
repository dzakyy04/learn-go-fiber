package route

import (
	"learn-go-fiber/config"
	"learn-go-fiber/handler"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	// Static asset
	r.Static("/public", config.ProjectRootPath+"/public")

	// Auth
	r.Post("/login", handler.LoginHandler)

	// User
	r.Get("/user", handler.UserHandlerGetAll)
	r.Get("/user/:id", handler.UserHandlerGetById)
	r.Post("/user", handler.UserHandlerCreate)
	r.Put("/user/:id", handler.UserHandlerUpdate)
	r.Put("/user/:id/email", handler.UserHandlerUpdateEmail)
	r.Delete("/user/:id", handler.UserHandlerDelete)
}
