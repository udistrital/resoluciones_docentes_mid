package models

// Respuesta es típica de una API que encapsula
// su respuesta en una propiedad "Data".
// NO USAR DIRECTAMENTE: La propiedad Data NO hace parte de
// este struct, usar Respuesta o RespuestaArr
type RespuestaAPI struct {
	Message string
	Status  string
	Success bool
}

// Respuesta es un RespuestaAPI donde
// los datos son un único objeto
type Respuesta struct {
	RespuestaAPI
	Data map[string]interface{}
}

// RespuestaModRes es un RespuestaAPI donde
// los datos son un único arreglo objeto
type RespuestaModRes struct {
	RespuestaAPI
	Data []*ModificacionResolucion
}

type RespuestaModVin struct {
	RespuestaAPI
	Data []*ModificacionVinculacion
}

type RespuestaVinculaciones struct {
	RespuestaAPI
	Data []*VinculacionDocente
}
