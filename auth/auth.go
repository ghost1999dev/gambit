package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type TokenJson struct {
	Sub       string
	Event_Id  string
	Token_use string
	Scope     string
	Auth_time int
	Iss       string
	Exp       int
	Iat       int
	Client_id string
	Username  string
}

func ValidoToken(token string) (bool, error, string) {
	//Partes del token
	//Cabecera
	//Payload
	//Firma

	//Separar el token
	splitToken := strings.Split(token, ".")
	if len(splitToken) !=3 {
		fmt.Println("Error al separar el token")
		return false, nil, "Error al separar el token"
	}
	//Decodificar el token
	userInfo , err :=base64.StdEncoding.DecodeString(splitToken[1])
	if err != nil {
		fmt.Println("Ha ocurrido un error al decodificar el token")
		return false,err,err.Error()
	}
	//Declarar una variable de tipo structura
	var tokenJson TokenJson
	//Declarar una variable para asignar los valores
	//base64 a la estructura
	err = json.Unmarshal(userInfo, &tokenJson)
	if err != nil {
		fmt.Println("No se puede decodificar la estructura json")
		return false,err,err.Error()
		
	}
	//Validar si el token expiro
	//Definir la fecha de ahora
	ahora := time.Now()
	fechaToken := time.Unix(int64(tokenJson.Exp),0)
	//Validar si la fecha del token esta antes de la fecha de ahora
	if fechaToken.Before(ahora) {
		fmt.Println("Fecha de expiracion de token = " + fechaToken.String())
		fmt.Println("Token expirado")
		return false, err,"Token expirado !!"
	}

	return true,nil, string(tokenJson.Username)

}