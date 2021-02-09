package helpers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/udistrital/resoluciones_docentes_mid/models"
)

func Buscar_Categoria_Docente(vigencia, periodo, documento_ident string) (categoria_nombre, categoria_id_old string, err error) {
	var temp map[string]interface{}
	var nombreCategoria string
	var idCategoriaOld string

	//TODO: quitar el hardconding para WSO2 cuando ya soporte https:
	err = GetJsonWSO2("http://"+beego.AppConfig.String("UrlcrudWSO2")+"/"+beego.AppConfig.String("NscrudUrano")+"/"+"categoria_docente/"+vigencia+"/"+periodo+"/"+documento_ident, &temp)
	if err != nil {
		return categoria_nombre, categoria_id_old, err
	}
	if temp != nil {

		jsonDocentes, err := json.Marshal(temp)

		if err != nil {
			return categoria_nombre, categoria_id_old, err
		}
		var tempDocentes models.ObjetoCategoriaDocente
		err = json.Unmarshal(jsonDocentes, &tempDocentes)
		if err != nil {
			return categoria_nombre, categoria_id_old, err
		}

		nombreCategoria = tempDocentes.CategoriaDocente.Categoria
		idCategoriaOld = tempDocentes.CategoriaDocente.IDCategoria

	}
	return nombreCategoria, idCategoriaOld, nil
}
