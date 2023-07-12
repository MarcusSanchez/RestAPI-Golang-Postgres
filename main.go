package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"log"
	"restAPI/database"
	"restAPI/router"
)

var port = "3000"

func main() {
	err := database.SetupPostgres()
	if err != nil {
		fmt.Println("Postgres connection failed")
		log.Fatal(err)
	}
	defer database.ClosePostgres()

	app := initializeFiber()
	router.StartRouting(app)
	startListening(app, port)
}

func initializeFiber() *fiber.App {
	app := fiber.New()

	// middleware
	app.Use(logger.New())
	app.Use(cors.New())
	app.Use(recover.New())

	//app.Use(parameters.New())
	return app
}

func startListening(app *fiber.App, port string) {
	log.Fatal(app.Listen(":" + port))
}
