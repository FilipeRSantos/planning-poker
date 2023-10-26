package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"liposo/planing-poker/adapters"
	"liposo/planing-poker/endpoints"
	"log"
)

func main() {
	_, err := adapters.CreateClient()
	if err != nil {
		panic(err)
	}

	engine := html.New("./templates", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	endpoints.SetupPortalEndpoints(app)
	endpoints.SetupApiEndpoints(app)

	log.Fatal(app.Listen(":8080"))
}
