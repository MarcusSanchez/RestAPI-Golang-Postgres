package helpers

import (
	"github.com/gofiber/fiber/v2"
	"restAPI/database"
	"restAPI/models"
	"restAPI/router/handlers/helpers/common"
)

func GetAllTeamPlayers__(teamName string) ([]models.Player, error) {
	team, err := common.FindTeam(teamName)
	if err != nil {
		return nil, err
	}

	var players []models.Player

	// Retrieve all players for the specified team from the players table
	if err = database.MyDB.Where("team_id = ?", team.ID).Find(&players).Error; err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to retrieve team players")
	}
	if len(players) == 0 {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "No players to return")
	}
	return players, nil
}

func CreatePlayer__(teamName string, c *fiber.Ctx) (*models.Player, error) {
	// Parse the JSON request body into a new Player object
	player := new(models.Player)
	if err := c.BodyParser(player); err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	if valid := common.ValidatePlayer(player); !valid {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Request body missing fields for player")
	}

	_, err := common.FindPlayer(teamName, player.PlayerName)
	if err == nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "This player already exists in the Database")
	}

	existingTeam, err := common.FindTeam(teamName)
	if err != nil {
		return nil, err
	}

	// Assign the team to the player
	player.TeamID = existingTeam.ID

	// Create the player in the database
	if err = database.MyDB.Create(player).Error; err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to create player")
	}

	return player, nil
}

func DeleteAllPlayers__(teamName string) error {
	team, err := common.FindTeam(teamName)
	if err != nil {
		return err
	}

	result := database.MyDB.Where("team_id = ?", team.ID).Delete(&models.Player{})
	if result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to delete players")
	}
	if result.RowsAffected == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "This team does not have any players")
	}
	return nil
}
