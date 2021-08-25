package helpers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/udistrital/resoluciones_docentes_mid/models"
)

func Buscar_Categoria_Docente(vigencia, periodo, documento_ident string) (categoria_nombre, categoria_id_old string, outputError map[string]interface{}) {
	var temp map[string]interface{}
	var nombreCategoria string
	var idCategoriaOld string

	//TODO: quitar el hardconding para WSO2 cuando ya soporte https:
	if response, err := GetJsonWSO2Test("http://"+beego.AppConfig.String("UrlcrudWSO2")+"/"+beego.AppConfig.String("NscrudUrano")+"/"+"categoria_docente/"+vigencia+"/"+periodo+"/"+documento_ident, &temp); response == 200 && err == nil {
		if temp != nil {

			jsonDocentes, err := json.Marshal(temp)

			if err != nil {
				outputError = map[string]interface{}{"funcion": "/Buscar_Categoria_Docente1", "err": "Error codificando la respuesta", "status": "404"}
				return categoria_nombre, categoria_id_old, outputError
			}
			var tempDocentes models.ObjetoCategoriaDocente
			err = json.Unmarshal(jsonDocentes, &tempDocentes)
			if err != nil {
				outputError = map[string]interface{}{"funcion": "/Buscar_Categoria_Docente2", "err": "Error decodificando la respuesta", "status": "404"}
				return categoria_nombre, categoria_id_old, outputError
			}

			nombreCategoria = tempDocentes.CategoriaDocente.Categoria
			idCategoriaOld = tempDocentes.CategoriaDocente.IDCategoria

		}
	} else {
		outputError = map[string]interface{}{"funcion": "/Buscar_Categoria_Docente3", "err": err.Error(), "status": "404"}
		return categoria_nombre, categoria_id_old, outputError
	}

	return nombreCategoria, idCategoriaOld, nil
}
