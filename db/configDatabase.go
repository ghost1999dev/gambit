package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/gambit/models"
	"github.com/gambit/secretmaneger"
	_ "github.com/go-sql-driver/mysql"
	
)
var SecretModel  models.SecretRDSJson
var err error
var Db *sql.DB

func ReadSecret() error {
	SecretModel, err = secretmaneger.GetSecret(os.Getenv("SecretName"))
	return err
}
func DbConnect() error {
	Db, err = sql.Open("mysql",ConnStr(SecretModel))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	err = Db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("Conexion exitosa de la base de datos")
	return nil
}
func ConnStr(keys models.SecretRDSJson) string{
	var dbUser, authToken, dbEnpoint,dbName string
	dbUser = keys.Username
	authToken = keys.Passwrod
	dbEnpoint = keys.Host
	dbName = "gambit"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPasswords=true",dbUser,authToken,dbEnpoint,dbName)
	fmt.Println(dsn)
	return dsn;
}
func UserIsAdmin(userUUID string)(bool,string){
	fmt.Println("Comienza UserIsAdmin")
	err:=DbConnect()
	if err !=nil {
		return false,err.Error()
	}
	defer Db.Close()
	sentencia := "SELECT 1 FROM users WHERE User_UUID='" + userUUID + "' AND User_Status = 0"
	fmt.Println(sentencia)
	rows,err :=Db.Query(sentencia)
	if err != nil {
		return false,err.Error()
	}
	var valor string

	//Se posiciona en el primer
	//econtro y sigue
	rows.Next()
	//Lee el valor y lo guarda en el apuntador
	rows.Scan(&valor)
	fmt.Println("UserIsAdmin > Ejecucion exitosa - valor devuelto" +valor)
	if valor == "1" {
		return true , ""
	}

	return false,"El usuario no es admin"
}