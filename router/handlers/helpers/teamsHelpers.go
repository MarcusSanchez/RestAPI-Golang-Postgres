package helpers

import (
	"github.com/gofiber/fiber/v2"
	"restAPI/database"
	"restAPI/models"
	"restAPI/router/handlers/helpers/common"
)

func GetAllTeams__() ([]models.Team, error) {
	// Retrieve all teams from the database
	var teams []models.Team
	if err := database.MyDB.Find(&teams).Error; err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to retrieve teams")
	}
	if len(teams) == 0 {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "No teams to return")
	}
	return teams, nil
}

func CreateTeam__(c *fiber.Ctx) (*models.Team, error) {
	team := new(models.Team)
	if err := c.BodyParser(team); err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	_, err := common.FindTeam(team.TeamName)
	if err == nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Team already Exists")
	}

	if valid := common.ValidateTeam(team); !valid {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Request body missing fields for team")
	}

	for _, player := range team.Players {
		if valid := common.ValidatePlayer(&player); !valid {
			return nil, fiber.NewError(fiber.StatusBadRequest, "Request body missing fields for player: "+player.PlayerName)
		}
	}

	// Save the team to the database
	if err = database.MyDB.Create(team).Error; err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to create team")
	}

	return team, nil
}

func DeleteAllTeams__() error {
	// Truncate the players table
	if err := database.MyDB.Exec("TRUNCATE TABLE players CASCADE").Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to clear players table")
	}

	// Truncate the teams table
	if err := database.MyDB.Exec("TRUNCATE TABLE teams CASCADE").Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to clear teams table")
	}

	return nil
}
