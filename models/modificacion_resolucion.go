package models

import (
	"time"
)

type ModificacionResolucion struct {
	Id                   int         
	ResolucionNuevaId    *Resolucion 
	ResolucionAnteriorId *Resolucion 
	Activo               bool        
	FechaCreacion        time.Time   
	FechaModificacion    time.Time   
}