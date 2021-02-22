package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mrrizal/rest-api-blog/interfaces"
	"github.com/mrrizal/rest-api-blog/models"
	"github.com/mrrizal/rest-api-blog/viewmodels"
	log "github.com/sirupsen/logrus"
)

type UserController struct {
	interfaces.IUserService
}

func (controller *UserController) GetUsersHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := models.UserModel{
			ID:       1,
			Username: "test_user",
			Email:    "test_user@gmail.com"}

		return c.JSON(viewmodels.UserVM{Username: user.Username, Email: user.Email})
	}
}

func (controller *UserController) SaveUserHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var user models.UserModel

		if err := c.BodyParser(&user); err != nil {
			log.Fatal(err)
		}

		user, err := controller.SaveUserService(user)
		if err != nil {
			msg := map[string]interface{}{
				"message": err.Error(),
			}
			return c.Status(400).JSON(msg)
		}
		return c.JSON(user)
	}
}
