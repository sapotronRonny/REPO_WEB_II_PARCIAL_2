package main

import (
	"net/http"

	"github.com/danie/APWII_PROYECTO/database"
	graphql "github.com/danie/APWII_PROYECTO/internal/graph" // Importa tu esquema
	"github.com/danie/APWII_PROYECTO/models"
	"github.com/gofiber/fiber/v2"
	"github.com/graphql-go/handler"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=kevinbd.postgres.database.azure.com user=administrador password=Picoreyes2003 dbname=apw2 port=5432 sslmode=prefer"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error al conectar a la base de datos: " + err.Error())
	}

	// Migrar la base de datos
	database.Migrate(db)

	// Inicializar la conexión a la base de datos en el paquete models
	models.InitDB(db)

	// Configurar Fiber
	app := fiber.New()

	// Configurar GraphQL handler
	h := handler.New(&handler.Config{
		Schema:   &graphql.Schema, // Usa tu esquema importado
		Pretty:   true,
		GraphiQL: true,
	})

	// GraphQL route
	app.Post("/graphql", func(c *fiber.Ctx) error {
		// Adaptar el request y response de Fiber a net/http
		fasthttpadaptor.NewFastHTTPHandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		})(c.Context())

		return nil
	})
	// GraphQL route
	app.Get("/graphql", func(c *fiber.Ctx) error {
		// Adaptar el request y response de Fiber a net/http
		fasthttpadaptor.NewFastHTTPHandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		})(c.Context())

		return nil
	})

	// Iniciar el servidor
	err = app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
