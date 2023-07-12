package handlers

import (
	"github.com/gofiber/fiber/v2"
	"restAPI/router/handlers/helpers"
)

func GetAllTeams(c *fiber.Ctx) error {
	teams, err := helpers.GetAllTeams__()
	if err != nil {
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(teams)
}

func CreateTeam(c *fiber.Ctx) error {
	team, err := helpers.CreateTeam__(c)
	if err != nil {
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(team)
}

func DeleteAllTeams(c *fiber.Ctx) error {
	err := helpers.DeleteAllTeams__()
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
