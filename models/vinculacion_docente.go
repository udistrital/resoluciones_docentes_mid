package models

import (
	"time"
)

type VinculacionDocente struct {
	Id                             int
	NumeroContrato                 string
	Vigencia                       int
	PersonaId                      int
	NumeroHorasSemanales           int
	NumeroSemanas                  int
	PuntoSalarialId                int
	SalarioMinimoId                int
	ResolucionVinculacionDocenteId *ResolucionVinculacionDocente
	DedicacionId                   *Dedicacion
	ProyectoCurricularId           int
	ValorContrato                  float64
	Categoria                      string
	Disponibilidad                 int
	DependenciaAcademica           int
	NumeroRp                       float64
	VigenciaRp                     float64
	FechaInicio                    time.Time
	Activo                         bool
	FechaCreacion                  string
	FechaModificacion              string
	VigenciaCarga                  int
	PeriodoCarga                   int
	NumeroHorasNuevas              int
	NumeroHorasModificacion        int
	NumeroSemanasNuevas            int
	NumeroSemanasRestantes         int
	FechaInicioNueva               time.Time
	Dedicacion                     string
	NivelAcademico                 string
	Periodo                        int
	NombreCompleto                 string
	NumeroDisponibilidad           int
	LugarExpedicionCedula          string
	TipoDocumento                  string
	NumeroMeses                    string
	ValorModificacionFormato       string
	ValorContratoFormato           string
	NumeroMesesNuevos              string
	ValorContratoInicialFormato    string
	ProyectoNombre                 string
}
