package handlers

import (
	"fmt"
	"strconv"

	//"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/gambit/auth"
	"github.com/gambit/routers"
)

func Manejadores(path string, method string,
	body string, header map[string]string,
	request events.APIGatewayV2HTTPRequest) (int,string){
		fmt.Println("Voy a procesar " +path+ ">"+method)
		id := request.PathParameters["id"]
		idn, _ :=strconv.Atoi(id)
		success,statusCode,user:=validoAutorizacion(path,method,header)
		if !success {
			return statusCode,user
		}
		switch path[0:4]{
		case "user":
			return ProcesoUsuario(body,path,method,user,id,request)
		case "prod":
			return ProcesoProducto(body,path,method,user,idn,request)
		case "stoc":
			return ProcesoStock(body,path,method,user,idn,request)
		case "addr":
			return ProcesoAddress(body,path,method,user,idn,request)
		case "cate":
			return ProcesoCategory(body,path,method,user,idn,request)
		case "orde":
			return ProcesoOrder(body,path,method,user,idn,request)
		}
		return 400, "Metodo invalido"
	}

func validoAutorizacion (path string, method string, headers map[string]string)(bool,int,string){
	if (path == "product" && method =="GET")||
	   (path == "category" && method == "GET"){
		return true,200,""	
	}
	token := headers["authorization"]
	if len(token)==0 {
		return false,401, "Token requerido"
	}
	exitoso,err,msg := auth.ValidoToken(token)
	if !exitoso {
		if err !=nil {
			fmt.Println("Error en el token " +err.Error())
			return false,401,err.Error()
			
		}else{
			fmt.Println("Error en el token " + msg)
			return false,401,err.Error()
		}
	}
	fmt.Println("Ok")
	return true,200,msg

}
func ProcesoUsuario (body string, path string,method string, user string, id string, request events.APIGatewayV2HTTPRequest) (int,string){
	return 400,"Metodo invalido"
}
func ProcesoProducto (body string, path string,method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int,string){
	return 400,"Metodo invalido"
}
func ProcesoStock (body string, path string, method string, user string,id int, request events.APIGatewayV2HTTPRequest) (int,string){
	return 400,"Metodo invalido"
}
func ProcesoAddress (body string, path string, method string, user string,id int, request events.APIGatewayV2HTTPRequest) (int,string){
	return 400,"Metodo invalido"
}
func ProcesoCategory (body string, path string, method string, user string,id int, request events.APIGatewayV2HTTPRequest) (int,string){
	switch method {
	case "GET":
		routers.InsertCategory(body,user)
	}
	return 400,"Metodo invalido"
}
func ProcesoOrder (body string, path string, method string, user string,id int, request events.APIGatewayV2HTTPRequest) (int,string){
	return 400,"Metodo invalido"
}