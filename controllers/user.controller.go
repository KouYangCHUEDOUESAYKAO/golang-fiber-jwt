package controllers

import (
	"github.com/KouYangCHUEDOUESAYKAO/golang-fiber-jwt/models"
	"github.com/gofiber/fiber/v2"
)

func GetMe(c *fiber.Ctx) error {
	user := c.Locals("user").(models.UserResponse)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "Get user success", "data": fiber.Map{"user": user}})
}
