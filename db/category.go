package db

import (
	"database/sql"
	"fmt"

	"github.com/gambit/models"
)

func InsertCategory(user models.Category)(int64,error){
	fmt.Println("Comienza registro...")
	err :=DbConnect()
	if err !=nil {
		return 0,err
	}
	defer Db.Close()
	sentencia := "INSERT INTO category (Categ_Name, Categ_Path) VALUES ('"+user.CategName + "','" +user.CategPath+"')"
	var result sql.Result
	result, err = Db.Exec(sentencia)
	if err !=nil {
		fmt.Println(err.Error())
		return 0,err
	}
	LastInsertId,err2 := result.LastInsertId()
	if err2 != nil {
		fmt.Println(err2.Error())
		return 0, err2
	}
	fmt.Println("Insert Category > Ejecusion exitosa")
	return LastInsertId,err2

}