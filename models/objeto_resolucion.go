package models

type ObjetoResolucion struct {
	Resolucion                   *Resolucion
	ResolucionVinculacionDocente *ResolucionVinculacionDocente
	ResolucionVieja              int
	NomDependencia               string
}
