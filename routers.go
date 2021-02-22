package main

import (
	"sync"

	"github.com/gofiber/fiber/v2"
)

type IFiberRouter interface {
	InitRouter() *fiber.App
}

type router struct{}

func (r *router) InitRouter() *fiber.App {

	app := fiber.New()

	api := app.Group("/api")
	v1 := api.Group("/v1")

	users := v1.Group("/users")
	userController := ServiceContainer().InjectUserController()
	users.Get("/", userController.GetUsersHandler())
	users.Post("/", userController.SaveUserHandler())

	return app
}

var (
	m          *router
	routerOnce sync.Once
)

func FiberRouter() IFiberRouter {
	if m == nil {
		routerOnce.Do(func() {
			m = &router{}
		})
	}
	return m
}
