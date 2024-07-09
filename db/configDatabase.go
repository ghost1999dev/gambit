package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gambit/models"
	"github.com/gambit/secretmaneger"
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