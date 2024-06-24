package routes

import (
	"github.com/danie/APWII_PROYECTO/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	// Rutas para Serie
	app.Get("/serie", handlers.GetSeries)
	app.Get("/serie/:id", handlers.GetSerieID)
	app.Post("/serie", handlers.CreateSerie)
	app.Put("/serie/:id", handlers.UpdateSerie)
	app.Delete("/serie/:id", handlers.DeleteSerie)

	// Rutas para Actor
	app.Get("/actor", handlers.GetActores)
	app.Get("/actor/:id", handlers.GetActorID)
	app.Post("/actor", handlers.CreateActor)
	app.Put("/actor/:id", handlers.UpdateActor)
	app.Delete("/actor/:id", handlers.DeleteActor)
}
