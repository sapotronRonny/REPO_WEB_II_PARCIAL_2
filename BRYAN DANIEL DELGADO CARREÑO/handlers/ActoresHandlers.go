package handlers

import (
	"github.com/danie/APWII_PROYECTO/models"
	"github.com/gofiber/fiber/v2"
)

// Controladores para Actor

func GetActores(c *fiber.Ctx) error {
	var actores []models.Actores
	models.DB.Find(&actores)
	return c.JSON(actores)
}

func GetActorID(c *fiber.Ctx) error {
	id := c.Params("id")
	var actores models.Actores
	models.DB.First(&actores, id)
	return c.JSON(actores)
}

func CreateActor(c *fiber.Ctx) error {
	var actores models.Actores
	if err := c.BodyParser(&actores); err != nil {
		return err
	}
	models.DB.Create(&actores)
	return c.JSON(actores)
}

func UpdateActor(c *fiber.Ctx) error {
	id := c.Params("id")
	var actores models.Actores
	models.DB.First(&actores, id)
	if err := c.BodyParser(&actores); err != nil {
		return err
	}
	models.DB.Save(&actores)
	return c.JSON(actores)
}

func DeleteActor(c *fiber.Ctx) error {
	id := c.Params("id")
	var actores models.Actores
	models.DB.Delete(&actores, id)
	return nil
}
