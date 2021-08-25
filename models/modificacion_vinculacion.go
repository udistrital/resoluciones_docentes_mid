package models

import (
	"time"
)

type ModificacionVinculacion struct {
	Id                             int                     
	ModificacionResolucionId       *ModificacionResolucion 
	VinculacionDocenteCanceladaId  *VinculacionDocente     
	VinculacionDocenteRegistradaId *VinculacionDocente     
	Horas                          int                 
	Activo                         bool                    
	FechaCreacion                  time.Time               
	FechaModificacion              time.Time               
}