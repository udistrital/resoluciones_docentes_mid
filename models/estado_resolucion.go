package models

import (
	"time"
)

type EstadoResolucion struct {
	Id                int       
	NombreEstado      string    
	Activo            bool      
	FechaCreacion     time.Time 
	FechaModificacion time.Time 
}