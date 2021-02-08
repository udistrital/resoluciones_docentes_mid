package models

type Ciudad struct {
	Id             int
	IdDepartamento int
	Nombre         string
	Abreviatura    string
	Descripcion    string
	Estado         string
	AbPais         string
	Departamento   string
	Poblacion      int
	Longitud       float64
	Latitud        float64
}
