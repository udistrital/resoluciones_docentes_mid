package models

import (
	"time"
)

type InformacionPersonaNatural struct {
	TipoDocumento                     *ParametroEstandar
	Id                                string
	DigitoVerificacion                float64
	PrimerApellido                    string
	SegundoApellido                   string
	PrimerNombre                      string
	SegundoNombre                     string
	Cargo                             string
	IdPaisNacimiento                  float64
	Perfil                            *ParametroEstandar
	Profesion                         string
	Especialidad                      string
	MontoCapitalAutorizado            float64
	Genero                            string
	GrupoEtnico                       string
	ComunidadLgbt                     bool
	CabezaFamilia                     bool
	PersonasACargo                    bool
	NumeroPersonasACargo              float64
	EstadoCivil                       string
	Discapacitado                     bool
	TipoDiscapacidad                  string
	DeclaranteRenta                   bool
	MedicinaPrepagada                 bool
	ValorUvtPrepagada                 float64
	CuentaAhorroAfc                   bool
	NumCuentaBancariaAfc              string
	IdEntidadBancariaAfc              float64
	InteresViviendaAfc                float64
	DependienteHijoMenorEdad          bool
	DependienteHijoMenos23Estudiando  bool
	DependienteHijoMas23Discapacitado bool
	DependienteConyuge                bool
	DependientePadreOHermano          bool
	IdNucleoBasico                    float64
	IdArl                             int
	IdEps                             int
	IdFondoPension                    int
	IdCajaCompensacion                int
	IdNitArl                          float64
	IdNitEps                          float64
	IdNitFondoPension                 float64
	IdNitCajaCompensacion             float64
	FechaExpedicionDocumento          time.Time
	IdCiudadExpedicionDocumento       float64
}
