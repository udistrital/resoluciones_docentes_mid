package models

type ModificacionVinculacion struct {
	Id                           int
	ModificacionResolucion       *ModificacionResolucion
	VinculacionDocenteCancelada  *VinculacionDocente
	VinculacionDocenteRegistrada *VinculacionDocente
	Horas                        int
}
