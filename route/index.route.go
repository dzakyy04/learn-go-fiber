package route

import (
	"learn-go-fiber/config"
	"learn-go-fiber/handler"
	"learn-go-fiber/middleware"
	"learn-go-fiber/utils"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(route *fiber.App) {
	// Static asset
	route.Static("/public", config.ProjectRootPath+"/public")

	// Auth
	route.Post("/login", handler.LoginHandler)
	route.Get("/me", middleware.AuthMiddleware, handler.GetUserDataHandler)

	// User
	route.Get("/user/", handler.UserHandlerGetAll)
	route.Get("/user/:id", handler.UserHandlerGetById)
	route.Post("/user/", handler.UserHandlerCreate)
	route.Put("/user/:id", handler.UserHandlerUpdate)
	route.Put("/user/:id/email", handler.UserHandlerUpdateEmail)
	route.Delete("/user/:id", handler.UserHandlerDelete)

	// Book
	route.Post("/book", utils.HandleSingleFile, handler.BookHandlerCreate)
}
