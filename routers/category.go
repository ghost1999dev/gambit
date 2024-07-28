package routers

import (
	"encoding/json"
	"strconv"

	"github.com/gambit/db"
	"github.com/gambit/models"
)

//Crear una funcion para insertar la categoria

func InsertCategory (body string, User string)(int, string){
	var t models.Category
	err := json.Unmarshal([]byte(body), &t)
	if err !=nil {
		return 400,"Error en los datos recibidos " +err.Error()
	}
	userIsAdmin,msg :=db.UserIsAdmin(User)
	if !userIsAdmin {
		return 400, msg 
	}
	result, err2 := db.InsertCategory(t)
	if err2 != nil {
		return 400, "Ocurrio un error al intentar realizar el registro de categoria " +t.CategName+">"+err2.Error()
	}
	return 200, "{CategID :"+ strconv.Itoa(int(result)) + "}"


}