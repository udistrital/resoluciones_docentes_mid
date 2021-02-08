package models

import (
	"time"
)

type Resolucion struct {
	Id                      int
	NumeroResolucion        string
	FechaExpedicion         time.Time
	Vigencia                int
	IdDependencia           int
	IdTipoResolucion        *TipoResolucion
	PreambuloResolucion     string
	ConsideracionResolucion string
	Estado                  bool
	FechaRegistro           time.Time
	Objeto                  string
	NumeroSemanas           int
	Periodo                 int
	Titulo                  string
	IdDependenciaFirma      int
	PeriodoCarga            int
	VigenciaCarga           int
}
