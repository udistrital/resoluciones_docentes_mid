package models

import (
	"time"
)

type ResolucionEstado struct {
	Id            int
	FechaRegistro time.Time
	Usuario       string
	Estado        *EstadoResolucion
	Resolucion    *Resolucion
}
