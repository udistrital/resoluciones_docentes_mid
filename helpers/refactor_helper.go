package helpers

import (
	"encoding/json"
	"fmt"
)

func LimpiezaRespuestaRefactor(respuesta map[string]interface{}, v interface{}) {
	//fmt.Println("CAMBIO RESPUESTA")
	//fmt.Println("respuesta")
	//fmt.Println(respuesta)
	/**
	* Cuando La petición no retorna datos, es decir, llega vacia no se carga la estructura
	**/
	datatype := fmt.Sprintf("%v", respuesta["Data"])
	//fmt.Println(datatype)
	switch datatype {
	case "map[]", "[map[]]": // response vacio
		break
	default:
		b, err := json.Marshal(respuesta["Data"])
		//fmt.Println(b)
		if err != nil {
			panic(err)
		}
		json.Unmarshal(b, v)
		respuesta = nil
	}
	//fmt.Println(v)
}
