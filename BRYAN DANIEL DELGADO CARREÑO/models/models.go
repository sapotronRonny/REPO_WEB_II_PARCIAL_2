package models

import (
	"time"

	"gorm.io/gorm"
)

var DB *gorm.DB // Exportamos la variable DB

// Función para inicializar la conexión con la base de datos
func InitDB(db *gorm.DB) {
	DB = db
}

type Serie struct {
	ID_Serie         uint `gorm:"primaryKey"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
	ActoresID        uint
	Actor            ActorSerie `gorm:"foreignKey:ActoresID;references:ID_Actor"`
	GeneroID         uint
	Genero           GeneroSerie `gorm:"foreignKey:GeneroID;references:ID_GeneroSerie"`
	Titulo           string      `gorm:"type:varchar(100)"`
	FechaCreacion    time.Time   `gorm:"type:date"`
	DuracionSegundos int64       // Guardar la duración en segundos
}

// Definir los estados permitidos
type Estado string

const (
	Activo   Estado = "Activo"
	Retirado Estado = "Retirado"
	Muerto   Estado = "Muerto"
)

// Estructura Actor
type ActorSerie struct {
	ID_Actor           uint `gorm:"primaryKey;unique"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          gorm.DeletedAt `gorm:"index"`
	Nombre             string         `gorm:"type:varchar(255)"`
	FechaInicioCarrera time.Time      `gorm:"type:date"`
	Estado             Estado         `gorm:"type:varchar(20)"`
}

type GeneroSerie struct {
	ID_GeneroSerie uint `gorm:"primaryKey;unique"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
	NombreGenero   string         `gorm:"type:varchar(255)"`
}
