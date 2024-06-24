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
	gorm.Model
	ID            uint `gorm:"primaryKey"`
	ActoresID     uint
	Actor         Actores       `gorm:"foreignKey:ActoresID"` // Agrega esta línea para la relación
	Titulo        string        `gorm:"type:varchar(100)"`
	FechaCreacion time.Time     `gorm:"type:date"`
	Duracion      time.Duration `gorm:"type:time"`
}

// Definir los estados permitidos
type Estado string

const (
	Activo   Estado = "Activo"
	Retirado Estado = "Retirado"
	Muerto   Estado = "Muerto"
)

// Estructura Actores
type Actores struct {
	gorm.Model
	Nombre             string    `gorm:"type:varchar(255)"`
	FechaInicioCarrera time.Time `gorm:"type:date"`
	Estado             Estado    `gorm:"type:varchar(20)"` // Usa el tipo Estado y define el tipo de columna en la base de datos
}
