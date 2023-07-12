package handlers

import (
	"github.com/gofiber/fiber/v2"
	"restAPI/router/handlers/helpers"
	"strings"
)

func GetTeam(c *fiber.Ctx) error {
	team, err := helpers.GetTeam__(strings.ReplaceAll(c.Params("teamName"), "-", " "))
	if err != nil {
		return c.JSON(fiber.Map{
			"error":        err.Error(),
		})
	}
	return c.JSON(team)
}

func UpdateTeam(c *fiber.Ctx) error {
	err := helpers.UpdateTeam__(strings.ReplaceAll(c.Params("teamName"), "-", " "), c)
	if err != nil {
		return c.JSON(fiber.Map{
			"acknowledged": "false",
			"error":        err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"acknowledged": "true",
	})
}

func DeleteTeam(c *fiber.Ctx) error {
	err := helpers.DeleteTeam__(strings.ReplaceAll(c.Params("teamName"), "-", " "))
	if err != nil {
		return c.JSON(fiber.Map{
			"acknowledged": "false",
			"error":        err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"acknowledged": "true",
	})
}
