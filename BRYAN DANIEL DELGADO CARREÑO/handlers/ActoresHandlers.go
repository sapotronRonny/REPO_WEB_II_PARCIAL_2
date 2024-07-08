package handlers

import (
	"github.com/danie/APWII_PROYECTO/models"
	"github.com/gofiber/fiber/v2"
)

// Controladores para Actor

func GetActorSerie(c *fiber.Ctx) error {
	var actores []models.ActorSerie
	models.DB.Find(&actores)
	return c.JSON(actores)
}

func GetActorSerieID(c *fiber.Ctx) error {
	id := c.Params("id")
	var actores models.ActorSerie
	models.DB.First(&actores, id)
	return c.JSON(actores)
}

func CreateActorSerie(c *fiber.Ctx) error {
	var actores models.ActorSerie
	if err := c.BodyParser(&actores); err != nil {
		return err
	}
	models.DB.Create(&actores)
	return c.JSON(actores)
}

func UpdateActorSerie(c *fiber.Ctx) error {
	id := c.Params("id")
	var actores models.ActorSerie
	models.DB.First(&actores, id)
	if err := c.BodyParser(&actores); err != nil {
		return err
	}
	models.DB.Save(&actores)
	return c.JSON(actores)
}

func DeleteActorSerie(c *fiber.Ctx) error {
	id := c.Params("id")
	var actores models.ActorSerie
	models.DB.Delete(&actores, id)
	return nil
}
