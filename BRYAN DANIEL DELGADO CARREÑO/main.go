package main

import (
	"github.com/danie/APWII_PROYECTO/database"
	"github.com/danie/APWII_PROYECTO/models"
	"github.com/danie/APWII_PROYECTO/routes"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Conectar a la base de datos
	database.ConnectDatabase()

	// Definir la cadena de conexión DSN
	dsn := "host=kevinbd.postgres.database.azure.com user=administrador password=Picoreyes2003 dbname=apw2 port=5432 sslmode=prefer "

	// Establecer la conexión a la base de datos
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error al conectar a la base de datos: " + err.Error())
	}

	//migrar la base de datos
	database.Migrate(db)

	// Inicializar la conexión a la base de datos en el paquete models
	models.InitDB(db)

	app := fiber.New()

	// Configurar las rutas
	routes.SetupRoutes(app)

	// Iniciar la aplicación
	app.Listen(":3000")
}
