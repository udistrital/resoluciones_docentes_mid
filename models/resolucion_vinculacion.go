package models

import (
	"time"
)

type ResolucionVinculacion struct {
	Id                 	int       
	Estado             	string    
	Numero             	string    
	Vigencia           	int       
	Facultad           	int       
	NivelAcademico     	string    
	Dedicacion         	string    
	FechaExpedicion    	time.Time 
	NumeroSemanas      	int       
	Periodo            	int       
	TipoResolucion     	string    
	IdDependenciaFirma 	int       
	PeriodoCarga       	int       
	VigenciaCarga      	int  
	FacultadNombre 		string  
	FacultadFirmaNombre	string   
}