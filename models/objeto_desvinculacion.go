package models

type Objeto_Desvinculacion struct {
	IdModificacionResolucion int
	IdNuevaResolucion        int
	DisponibilidadNueva      DisponibilidadApropiacion
	TipoDesvinculacion       string
	DocentesDesvincular      []VinculacionDocente
}
