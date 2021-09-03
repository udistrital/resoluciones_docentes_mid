package models

type TipoResolucion struct {
	Id                   int
	NombreTipoResolucion string
	Descripcion          string
	Activo               bool
	FechaCreacion        string
	FechaModificacion    string
}
