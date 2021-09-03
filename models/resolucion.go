package models

import (
	"time"
)

type Resolucion struct {
	Id                      int
	NumeroResolucion        string
	FechaExpedicion         time.Time
	Vigencia                int
	DependenciaId           int
	TipoResolucionId        *TipoResolucion
	PreambuloResolucion     string
	ConsideracionResolucion string
	NumeroSemanas           int
	Periodo                 int
	Titulo                  string
	DependenciaFirmaId      int
	VigenciaCarga           int
	PeriodoCarga            int
	Activo                  bool
	FechaCreacion           string
	FechaModificacion       string
}
