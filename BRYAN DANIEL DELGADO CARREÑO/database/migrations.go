package database

import (
	"github.com/danie/APWII_PROYECTO/models"

	"gorm.io/gorm"
)

// Migrar modelos
func Migrate(db *gorm.DB) error {

	// Luego migra los modelos que dependen de este tipo
	if err := db.AutoMigrate(&models.Actores{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&models.Serie{}); err != nil {
		return err
	}

	return nil
}
