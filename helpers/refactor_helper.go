package helpers

import (
	"encoding/json"
	_ "fmt"
)

func LimpiezaRespuestaRefactor(respuesta map[string]interface{}, v interface{}) {
	//fmt.Println("CAMBIO RESPUESTA")
	//fmt.Println("respuesta")
	//fmt.Println(respuesta)
	b, err := json.Marshal(respuesta["Data"])
	//fmt.Println(b)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(b, v)
	//fmt.Println(v)
}

