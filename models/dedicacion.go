package models

import (
	"time"
)

type Dedicacion struct {
	Id                int       
	NombreDedicacion  string    
	Descripcion       string    
	Activo            bool      
	FechaCreacion     time.Time 
	FechaModificacion time.Time 
}