package router

import (
	"github.com/gofiber/fiber/v2"
	"restAPI/router/handlers"
)

func StartRouting(app *fiber.App) {
	routeTeams(app)
	routeSingleTeam(app)
	routeSingleTeamPlayers(app)
	routeSinglePlayer(app)
}

func routeTeams(app *fiber.App) {
	teams := app.Group("/teams")
	teams.Get("/", handlers.GetAllTeams)
	teams.Post("/", handlers.CreateTeam)
	teams.Delete("/", handlers.DeleteAllTeams)
}

func routeSingleTeam(app *fiber.App) {
	team := app.Group("/teams/:teamName")
	team.Get("/", handlers.GetTeam)
	team.Put("/", handlers.UpdateTeam)
	team.Delete("/", handlers.DeleteTeam)
}

func routeSingleTeamPlayers(app *fiber.App) {
	teamPlayers := app.Group("/teams/:teamName/players")
	teamPlayers.Get("/", handlers.GetAllTeamPlayers)
	teamPlayers.Post("/", handlers.CreatePlayer)
	teamPlayers.Delete("/", handlers.DeleteAllPlayers)
}

func routeSinglePlayer(app *fiber.App) {
	player := app.Group("/teams/:teamName/players/:playerName")
	player.Get("/", handlers.GetPlayer)
	player.Put("/", handlers.UpdatePlayer)
	player.Delete("/", handlers.DeletePlayer)
}
