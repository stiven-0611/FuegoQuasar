// main package para el servidor Topsecret.
package main

import (
	"prueba/conection"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(conection.TopSecretPost)
}
