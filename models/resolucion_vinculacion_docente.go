package models

import (
	"time"
)

type ResolucionVinculacionDocente struct {
	Id                int      
	FacultadId        int      
	Dedicacion        string   
	NivelAcademico    string   
	Activo            bool     
	FechaCreacion     time.Time
	FechaModificacion time.Time
}
