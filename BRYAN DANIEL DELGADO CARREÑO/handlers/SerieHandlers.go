package handlers

import (
	"github.com/danie/APWII_PROYECTO/models"
	"github.com/gofiber/fiber/v2"
)

// Controladores para pelicula

func GetSeries(c *fiber.Ctx) error {
	var serie []models.Serie
	models.DB.Find(&serie)
	return c.JSON(serie)
}

func GetSerieID(c *fiber.Ctx) error {
	id := c.Params("id")
	var serie models.Serie
	models.DB.First(&serie, id)
	return c.JSON(serie)
}

func CreateSerie(c *fiber.Ctx) error {
	var serie models.Serie
	if err := c.BodyParser(&serie); err != nil {
		return err
	}
	models.DB.Create(&serie)
	return c.JSON(serie)
}

func UpdateSerie(c *fiber.Ctx) error {
	id := c.Params("id")
	var serie models.Serie
	models.DB.First(&serie, id)
	if err := c.BodyParser(&serie); err != nil {
		return err
	}
	models.DB.Save(&serie)
	return c.JSON(serie)
}

func DeleteSerie(c *fiber.Ctx) error {
	id := c.Params("id")
	var serie models.Serie
	models.DB.Delete(&serie, id)
	return nil
}
