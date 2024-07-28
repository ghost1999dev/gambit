package gambit

import (
	"context"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/gambit/awsgo"
	"github.com/gambit/db"
	"github.com/gambit/handlers"
)
func main(){
	lambda.Start(EjecutoLambda)
}
func EjecutoLambda(ctx context.Context, request events.APIGatewayV2HTTPRequest)(*events.APIGatewayProxyResponse, error){
	//Conectar con el servicio de aws
	awsgo.StartAWS()
	if !ValidoParamtros() {
		panic("Error en los parametros")
		
	}
	var res *events.APIGatewayProxyResponse
	prefijo := os.Getenv("UrlPrefix")
	path:= strings.Replace(request.RawPath, prefijo,"",-1)
	metodo := request.RequestContext.HTTP.Method
	body := request.Body
	header := request.Headers

	db.ReadSecret()

	//Llamar a los manejadores
	status,message := handlers.Manejadores(path,metodo,body,header,request)
	headersResp := map[string]string{
		"Content-Type": "application/json",
	}
	res = &events.APIGatewayProxyResponse{
		StatusCode: status,
		Body: string(message),
		Headers: headersResp,
	}
	return res,nil
	

}

func ValidoParamtros() bool{
	_, traeParametro := os.LookupEnv("SecretName")
	if !traeParametro {
		return traeParametro;
	}
	_,traeParametro = os.LookupEnv("UrlPrefix")
	if !traeParametro {
		return traeParametro
	}
	return traeParametro
}