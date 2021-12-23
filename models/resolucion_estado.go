package models

type ResolucionEstado struct {
	Id                 int
	Usuario            string
	EstadoResolucionId *EstadoResolucion
	ResolucionId       *Resolucion
	Activo             bool
	FechaCreacion      string
	FechaModificacion  string
}
