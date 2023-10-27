package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"liposo/planing-poker/endpoints"
	"liposo/planing-poker/handlers"
	"log"
)

func main() {
	apiHandler, err := handlers.CreateClient()
	if err != nil {
		panic(err)
	}

	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	endpoints.SetupPortalEndpoints(app, apiHandler)
	endpoints.SetupApiEndpoints(app, apiHandler)

	log.Fatal(app.Listen(":8080"))
}
