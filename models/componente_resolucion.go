package models

import (
	"time"
)

type ComponenteResolucion struct {
	Id                        int                   
	Numero                    int                   
	ResolucionId              *Resolucion           
	Texto                     string                
	TipoComponente            string                
	ComponenteResolucionPadre *ComponenteResolucion 
	Activo                    bool                  
	FechaCreacion             time.Time             
	FechaModificacion         time.Time             
}
