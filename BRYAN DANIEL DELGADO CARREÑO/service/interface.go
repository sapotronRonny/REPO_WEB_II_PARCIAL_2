package service

import "github.com/danie/APWII_PROYECTO/models"

// ActorService define las operaciones disponibles para la entidad ActorSerie.
type ActorService interface {
	GetActores() ([]models.ActorSerie, error)
	GetActorById(id string) (*models.ActorSerie, error)
	CreateActor(input map[string]interface{}) (*models.ActorSerie, error)
	UpdateActor(id string, input map[string]interface{}) (*models.ActorSerie, error)
	DeleteActor(id string) (bool, error)
}

// GeneroService define las operaciones disponibles para la entidad GeneroSerie.
type GeneroService interface {
	GetGeneros() ([]models.GeneroSerie, error)
	GetGeneroById(id string) (*models.GeneroSerie, error)
	CreateGenero(input map[string]interface{}) (*models.GeneroSerie, error)
	UpdateGenero(id string, input map[string]interface{}) (*models.GeneroSerie, error)
	DeleteGenero(id string) (bool, error)
}

// SerieService define las operaciones disponibles para la entidad Serie.
type SerieService interface {
	GetSeries() ([]models.Serie, error)
	GetSerieById(id string) (*models.Serie, error)
	CreateSerie(input map[string]interface{}) (*models.Serie, error)
	UpdateSerie(id string, input map[string]interface{}) (*models.Serie, error)
	DeleteSerie(id string) (bool, error)
}
