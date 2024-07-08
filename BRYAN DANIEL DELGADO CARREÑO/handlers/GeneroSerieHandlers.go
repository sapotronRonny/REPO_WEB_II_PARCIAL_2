package handlers

import (
	"github.com/danie/APWII_PROYECTO/models"
	"github.com/gofiber/fiber/v2"
)

// Controladores para pelicula

func GetGeneroSeries(c *fiber.Ctx) error {
	var Generoserie []models.GeneroSerie
	models.DB.Find(&Generoserie)
	return c.JSON(Generoserie)
}

func GetGeneroSerieID(c *fiber.Ctx) error {
	id := c.Params("id")
	var Generoserie models.GeneroSerie
	models.DB.First(&Generoserie, id)
	return c.JSON(Generoserie)
}

func CreateGeneroSerie(c *fiber.Ctx) error {
	var Generoserie models.GeneroSerie
	if err := c.BodyParser(&Generoserie); err != nil {
		return err
	}
	models.DB.Create(&Generoserie)
	return c.JSON(Generoserie)
}

func UpdateGeneroSerie(c *fiber.Ctx) error {
	id := c.Params("id")
	var Generoserie models.GeneroSerie
	models.DB.First(&Generoserie, id)
	if err := c.BodyParser(&Generoserie); err != nil {
		return err
	}
	models.DB.Save(&Generoserie)
	return c.JSON(Generoserie)
}

func DeleteGeneroSerie(c *fiber.Ctx) error {
	id := c.Params("id")
	var Generoserie models.GeneroSerie
	models.DB.Delete(&Generoserie, id)
	return nil
}
