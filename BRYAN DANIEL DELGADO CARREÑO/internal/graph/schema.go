package graphql

import (
	"time"

	"github.com/danie/APWII_PROYECTO/models"
	"github.com/graphql-go/graphql"
)

var (
	// Define el tipo Serie
	serieType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Serie",
		Fields: graphql.Fields{
			"ID_Serie": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"Titulo": &graphql.Field{
				Type: graphql.String,
			},
			"FechaCreacion": &graphql.Field{
				Type: graphql.DateTime,
			},
			"DuracionSegundos": &graphql.Field{
				Type: graphql.Int,
			},
			"Actor": &graphql.Field{
				Type: actorSerieType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					serie, _ := p.Source.(models.Serie)
					var actor models.ActorSerie
					models.DB.First(&actor, serie.ActoresID)
					return actor, nil
				},
			},
			"Genero": &graphql.Field{
				Type: generoSerieType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					serie, _ := p.Source.(models.Serie)
					var genero models.GeneroSerie
					models.DB.First(&genero, serie.GeneroID)
					return genero, nil
				},
			},
			// Define otros campos según tu modelo Serie
		},
	})

	// Define el tipo ActorSerie
	actorSerieType = graphql.NewObject(graphql.ObjectConfig{
		Name: "ActorSerie",
		Fields: graphql.Fields{
			"ID_Actor": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"Nombre": &graphql.Field{
				Type: graphql.String,
			},
			"FechaInicioCarrera": &graphql.Field{
				Type: graphql.DateTime,
			},
			"Estado": &graphql.Field{
				Type: graphql.String,
			},
			// Define otros campos según tu modelo ActorSerie
		},
	})

	// Define el tipo GeneroSerie
	generoSerieType = graphql.NewObject(graphql.ObjectConfig{
		Name: "GeneroSerie",
		Fields: graphql.Fields{
			"ID_GeneroSerie": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"NombreGenero": &graphql.Field{
				Type: graphql.String,
			},
			// Define otros campos según tu modelo GeneroSerie
		},
	})

	// Define el tipo Query
	rootQuery = graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"series": &graphql.Field{
				Type: graphql.NewList(serieType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					var series []models.Serie
					models.DB.Find(&series)
					return series, nil
				},
			},
			"actorSerie": &graphql.Field{
				Type: actorSerieType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, ok := p.Args["id"].(int)
					if !ok {
						return nil, nil
					}
					var actor models.ActorSerie
					models.DB.First(&actor, id)
					return actor, nil
				},
			},
			"generoSerie": &graphql.Field{
				Type: generoSerieType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, ok := p.Args["id"].(int)
					if !ok {
						return nil, nil
					}
					var genero models.GeneroSerie
					models.DB.First(&genero, id)
					return genero, nil
				},
			},
		},
	})

	// Define el tipo Mutation
	rootMutation = graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createSerie": &graphql.Field{
				Type: serieType,
				Args: graphql.FieldConfigArgument{
					"Titulo": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"FechaCreacion": &graphql.ArgumentConfig{
						Type: graphql.DateTime,
					},
					"DuracionSegundos": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"ActoresID": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"GeneroID": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					var duracionSegundos int64
					if duracion, ok := p.Args["DuracionSegundos"].(int); ok {
						duracionSegundos = int64(duracion)
					}
					fechaCreacion, _ := p.Args["FechaCreacion"].(time.Time)
					serie := models.Serie{
						Titulo:           p.Args["Titulo"].(string),
						FechaCreacion:    fechaCreacion,
						DuracionSegundos: duracionSegundos,
						ActoresID:        uint(p.Args["ActoresID"].(int)),
						GeneroID:         uint(p.Args["GeneroID"].(int)),
					}
					models.DB.Create(&serie)
					return serie, nil
				},
			},
			"updateSerie": &graphql.Field{
				Type: serieType,
				Args: graphql.FieldConfigArgument{
					"ID_Serie": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"Titulo": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"FechaCreacion": &graphql.ArgumentConfig{
						Type: graphql.DateTime,
					},
					"DuracionSegundos": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"ActoresID": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"GeneroID": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					var serie models.Serie
					models.DB.First(&serie, p.Args["ID_Serie"])
					if serie.ID_Serie == 0 {
						return nil, nil
					}
					if titulo, ok := p.Args["Titulo"].(string); ok {
						serie.Titulo = titulo
					}
					if fecha, ok := p.Args["FechaCreacion"].(time.Time); ok {
						serie.FechaCreacion = fecha
					}
					if duracion, ok := p.Args["DuracionSegundos"].(int); ok {
						serie.DuracionSegundos = int64(duracion)
					}
					if actoresID, ok := p.Args["ActoresID"].(int); ok {
						serie.ActoresID = uint(actoresID)
					}
					if generoID, ok := p.Args["GeneroID"].(int); ok {
						serie.GeneroID = uint(generoID)
					}
					models.DB.Save(&serie)
					return serie, nil
				},
			},
			"deleteSerie": &graphql.Field{
				Type: serieType,
				Args: graphql.FieldConfigArgument{
					"ID_Serie": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					var serie models.Serie
					models.DB.First(&serie, p.Args["ID_Serie"])
					if serie.ID_Serie == 0 {
						return nil, nil
					}
					models.DB.Delete(&serie)
					return serie, nil
				},
			},
			"createActorSerie": &graphql.Field{
				Type: actorSerieType,
				Args: graphql.FieldConfigArgument{
					"Nombre": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"FechaInicioCarrera": &graphql.ArgumentConfig{
						Type: graphql.DateTime,
					},
					"Estado": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					fechaInicio, _ := p.Args["FechaInicioCarrera"].(time.Time)
					actor := models.ActorSerie{
						Nombre:             p.Args["Nombre"].(string),
						FechaInicioCarrera: fechaInicio,
						Estado:             models.Estado(p.Args["Estado"].(string)),
					}
					models.DB.Create(&actor)
					return actor, nil
				},
			},
			"createGeneroSerie": &graphql.Field{
				Type: generoSerieType,
				Args: graphql.FieldConfigArgument{
					"NombreGenero": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					genero := models.GeneroSerie{
						NombreGenero: p.Args["NombreGenero"].(string),
					}
					models.DB.Create(&genero)
					return genero, nil
				},
			},
		},
	})

	// Define el esquema GraphQL
	Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	})
)
