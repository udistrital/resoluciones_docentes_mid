package models

type ModificacionVinculacion struct {
	Id                             int
	ModificacionResolucionId       *ModificacionResolucion
	VinculacionDocenteCanceladaId  *VinculacionDocente
	VinculacionDocenteRegistradaId *VinculacionDocente
	Horas                          int
	Activo                         bool
	FechaCreacion                  string
	FechaModificacion              string
}
