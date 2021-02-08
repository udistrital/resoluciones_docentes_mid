package models

import (
	"time"
)

type ResolucionVinculacion struct {
	Id                  int
	Estado              string
	Numero              string
	Vigencia            int
	Facultad            int
	NivelAcademico      string
	Dedicacion          string
	FechaExpedicion     time.Time
	NumeroSemanas       int
	Periodo             int
	TipoResolucion      string
	FacultadNombre      string
	IdDependenciaFirma  int
	FacultadFirmaNombre string
	VigenciaCarga       int
	PeriodoCarga        int
}