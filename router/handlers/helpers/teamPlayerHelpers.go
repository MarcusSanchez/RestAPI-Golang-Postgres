package helpers

import (
	"github.com/gofiber/fiber/v2"
	"restAPI/database"
	"restAPI/models"
	"restAPI/router/handlers/helpers/common"
)

func GetPlayer__(teamName, playerName string) (models.Player, error) {
	var player models.Player

	team, err := common.FindTeam(teamName)
	if err != nil {
		return player, err
	}

	// Retrieve the player from the players table
	if err = database.MyDB.Where("team_id = ? AND player_name = ?", team.ID, playerName).Take(&player).Error; err != nil {
		return player, fiber.NewError(fiber.StatusNotFound, "Player not found")
	}

	// Return the player to be marshalled into JSON
	return player, nil
}

func UpdatePlayer__(teamName, playerName string, c *fiber.Ctx) (*models.Player, error) {

	team, err := common.FindTeam(teamName)
	if err != nil {
		return nil, err
	}

	// Retrieve the existing player from the players table
	existingPlayer := new(models.Player)
	if err = database.MyDB.Where("team_id = ? AND player_name = ?", team.ID, playerName).Take(existingPlayer).Error; err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "Player not found")
	}

	// Parse the JSON request body into a new Player object
	updatedPlayer := new(models.Player)
	if err = c.BodyParser(updatedPlayer); err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	if updatedPlayer.PlayerName != "" {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Cannot update player name, only position")
	} else if updatedPlayer.Position != "" {
		existingPlayer.Position = updatedPlayer.Position
	} else {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Empty request body")
	}

	// Save the updated player in the database
	if err = database.MyDB.Save(existingPlayer).Error; err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to update player")
	}

	// Return the updated player as a JSON response
	return existingPlayer, nil
}

func DeletePlayer__(teamName, playerName string) error {
	team, err := common.FindTeam(teamName)
	if err != nil {
		return err
	}

	// Delete the player from the players table
	result := database.MyDB.Where("team_id = ? AND player_name = ?", team.ID, playerName).Delete(&models.Player{})
	if result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to delete player")
	}
	// Check the number of rows affected by the delete operation
	if result.RowsAffected == 0 {
		return fiber.NewError(fiber.StatusNotFound, "Player not found")
	}

	return nil
}
