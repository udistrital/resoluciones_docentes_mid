package models

import (
	"time"
)

type EstadoResolucion struct {
	Id            int
	FechaRegistro time.Time
	NombreEstado  string
}
