package models

type ComponenteResolucion struct {
	Id              int
	Numero          int
	ResolucionId    *Resolucion
	Texto           string
	TipoComponente  string
	ComponentePadre *ComponenteResolucion
}
