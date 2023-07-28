package controllers

import (
	"github.com/KouYangCHUEDOUESAYKAO/golang-fiber-jwt/models"
	"github.com/gofiber/fiber/v2"
)

func SignUpUser(c *fiber.Ctx) error {
	var payload *models.SignUpInput

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
}
