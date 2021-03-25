package models

import (
	"time"
)

type TipoResolucion struct {
	Id                   int       
	NombreTipoResolucion string    
	Descripcion          string    
	Activo               bool      
	FechaCreacion        time.Time 
	FechaModificacion    time.Time 
}