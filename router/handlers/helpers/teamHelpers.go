package helpers

import (
	"github.com/gofiber/fiber/v2"
	"restAPI/database"
	"restAPI/models"
	"restAPI/router/handlers/helpers/common"
)

func GetTeam__(teamName string) (*models.Team, error) {
	team, err := common.FindTeam(teamName)
	if err != nil {
		return &team, err
	}
	return &team, nil
}

func UpdateTeam__(teamName string, c *fiber.Ctx) error {
	existingTeam := new(models.Team)
	if err := database.MyDB.Where("team_name = ?", teamName).First(existingTeam).Error; err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Team not found")
	}

	updatedTeam := new(models.Team)
	if err := c.BodyParser(updatedTeam); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	if updatedTeam.TeamName == "" && updatedTeam.Location == "" && updatedTeam.Nickname == "" && len(updatedTeam.Players) == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "No updates to be made")
	}

	if updatedTeam.TeamName != "" && updatedTeam.TeamName != existingTeam.TeamName {
		_, err := common.FindTeam(updatedTeam.TeamName)
		if err == nil {
			return fiber.NewError(fiber.StatusBadRequest, "Team already exists")
		}
		existingTeam.TeamName = updatedTeam.TeamName
	}

	if updatedTeam.Location != "" && updatedTeam.Location != existingTeam.Location {
		existingTeam.Location = updatedTeam.Location
	}

	if updatedTeam.Nickname != "" && updatedTeam.Nickname != existingTeam.Nickname {
		existingTeam.Nickname = updatedTeam.Nickname
	}

	if len(updatedTeam.Players) != 0 {
		team, err := common.FindTeam(teamName)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Error updating players")
		}

		err = database.MyDB.Where("team_id = ?", team.ID).Delete(&models.Player{}).Error
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Error updating players")
		}
		existingTeam.Players = updatedTeam.Players
	}

	err := database.MyDB.Save(existingTeam).Error
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return nil
}

func DeleteTeam__(teamName string) error {
	team, err := common.FindTeam(teamName)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Team does not exist")
	}

	err = database.MyDB.Where("team_id = ?", team.ID).Delete(&models.Player{}).Error
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	result := database.MyDB.Where("team_name = ?", teamName).Delete(&models.Team{})
	if result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError, result.Error.Error())
	}
	return nil
}
