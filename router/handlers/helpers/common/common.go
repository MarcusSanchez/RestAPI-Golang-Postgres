package common

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"restAPI/database"
	"restAPI/models"
)

func FindTeam(teamName string) (models.Team, error) {
	var team models.Team

	// Retrieve the team from the database based on the team name
	if err := database.MyDB.Where("team_name = ?", teamName).First(&team).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return team, fiber.NewError(fiber.StatusNotFound, "Team not found")
		}
		return team, fiber.NewError(fiber.StatusInternalServerError, "Failed to retrieve team")
	}

	return team, nil
}

func FindPlayer(teamName string, playerName string) (models.Player, error) {
	var player models.Player

	team, err := FindTeam(teamName)
	if err != nil {
		return player, err
	}

	// Retrieve the player from the database based on the player name
	if err = database.MyDB.Where("player_name = ? AND team_id = ?", playerName, team.ID).First(&player).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return player, fiber.NewError(fiber.StatusNotFound, "Player not found")
		}
		return player, fiber.NewError(fiber.StatusInternalServerError, "Failed to retrieve player")
	}

	return player, nil
}

func ValidatePlayer(player *models.Player) bool {
	if player.PlayerName == "" || player.Position == "" {
		return false
	}
	return true
}

func ValidateTeam(team *models.Team) bool {
	if team.TeamName == "" || team.Location == "" || team.Nickname == "" {
		return false
	}
	return true
}
