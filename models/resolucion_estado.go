package models

import (
	"time"
)

type ResolucionEstado struct {
	Id                 int               
	Usuario            string            
	EstadoResolucionId *EstadoResolucion 
	ResolucionId       *Resolucion       
	Activo             bool              
	FechaCreacion      time.Time         
	FechaModificacion  time.Time         
}
