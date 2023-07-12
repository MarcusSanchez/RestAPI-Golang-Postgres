package handlers

import (
	"github.com/gofiber/fiber/v2"
	"restAPI/router/handlers/helpers"
	"strings"
)

func GetPlayer(c *fiber.Ctx) error {
	teamName := strings.ReplaceAll(c.Params("teamName"), "-", " ")
	playerName := strings.ReplaceAll(c.Params("playerName"), "-", " ")

	player, err := helpers.GetPlayer__(teamName, playerName)
	if err != nil {
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(player)
}

func UpdatePlayer(c *fiber.Ctx) error {
	teamName := strings.ReplaceAll(c.Params("teamName"), "-", " ")
	playerName := strings.ReplaceAll(c.Params("playerName"), "-", " ")

	player, err := helpers.UpdatePlayer__(teamName, playerName, c)
	if err != nil {
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(player)
}

func DeletePlayer(c *fiber.Ctx) error {
	// Get the teamName and playerName parameters from the request
	teamName := strings.ReplaceAll(c.Params("teamName"), "-", " ")
	playerName := strings.ReplaceAll(c.Params("playerName"), "-", " ")

	err := helpers.DeletePlayer__(teamName, playerName)
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
