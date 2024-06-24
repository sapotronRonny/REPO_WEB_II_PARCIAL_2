package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Parámetros de conexión a la base de datos
	dsn := "host=kevinbd.postgres.database.azure.com user=administrador password=Picoreyes2003 dbname=apw2 port=5432 sslmode=prefer "
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error al conectar a la base de datos: ", err)
	}

	DB = db

	log.Println("Conexión a la base de datos establecida exitosamente")
}
