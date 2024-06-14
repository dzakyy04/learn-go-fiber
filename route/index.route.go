package route

import (
	"learn-go-fiber/config"
	"learn-go-fiber/handler"
	"learn-go-fiber/middleware"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(route *fiber.App) {
	// Static asset
	route.Static("/public", config.ProjectRootPath+"/public")

	// Auth
	route.Post("/login", handler.LoginHandler)

	// User
	userGroup := route.Group("/user", middleware.AuthMiddleware)
	userGroup.Get("/", handler.UserHandlerGetAll)
	userGroup.Get("/:id", handler.UserHandlerGetById)
	userGroup.Post("/", handler.UserHandlerCreate)
	userGroup.Put("/:id", handler.UserHandlerUpdate)
	userGroup.Put("/:id/email", handler.UserHandlerUpdateEmail)
	userGroup.Delete("/:id", handler.UserHandlerDelete)
}
