package controllers

import (
	"myproject/models"

	"myproject/database"

	"github.com/gofiber/fiber/v2"
)

func RegisterUser(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: data["password"],
	}
	database.DB.Create(&user)

	// user := new(models.User)
	// if err := c.BodyParser(user); err != nil {
	// 	return err
	// }
	// database.DB.Create(&user)
	return c.JSON(&user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User
	database.DB.Where("email=?", data["email"]).First(&user)

	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "User Not Found :(",
		})
	}

	if data["password"] != user.Password {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "OOPS!! Incorrect Password! Try again.. :(",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Successful Login!!",
	})
}
