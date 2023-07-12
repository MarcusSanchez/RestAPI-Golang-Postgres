package handlers

import (
	"github.com/gofiber/fiber/v2"
	"restAPI/router/handlers/helpers"
	"strings"
)

func GetAllTeamPlayers(c *fiber.Ctx) error {
	players, err := helpers.GetAllTeamPlayers__(strings.ReplaceAll(c.Params("teamName"), "-", " "))
	if err != nil {
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(players)
}

func CreatePlayer(c *fiber.Ctx) error {
	player, err := helpers.CreatePlayer__(strings.ReplaceAll(c.Params("teamName"), "-", " "), c)
	if err != nil {
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(player)
}

func DeleteAllPlayers(c *fiber.Ctx) error {
	err := helpers.DeleteAllPlayers__(strings.ReplaceAll(c.Params("teamName"), "-", " "))
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
