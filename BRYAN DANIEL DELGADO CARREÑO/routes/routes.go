package routes

import (
	"github.com/danie/APWII_PROYECTO/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the API")
	})

	// Rutas para Serie
	app.Get("/serie", handlers.GetSeries)
	app.Get("/serie/:id", handlers.GetSerieID)
	app.Post("/serie", handlers.CreateSerie)
	app.Put("/serie/:id", handlers.UpdateSerie)
	app.Delete("/serie/:id", handlers.DeleteSerie)

	// Rutas para Actor
	app.Get("/actor", handlers.GetActorSerie)
	app.Get("/actor/:id", handlers.GetActorSerie)
	app.Post("/actor", handlers.CreateActorSerie)
	app.Put("/actor/:id", handlers.UpdateActorSerie)
	app.Delete("/actor/:id", handlers.DeleteActorSerie)

	// Rutas para Genero
	app.Get("/genero", handlers.GetGeneroSeries)
	app.Get("/genero/:id", handlers.GetGeneroSerieID)
	app.Post("/genero", handlers.CreateGeneroSerie)
	app.Put("/genero/:id", handlers.UpdateGeneroSerie)
	app.Delete("/genero/:id", handlers.DeleteGeneroSerie)
}
