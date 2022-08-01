package conection

import (
	"encoding/json"
	"fmt"
	"net/http"
	"prueba/topsecret"

	"github.com/aws/aws-lambda-go/events"
)

type H map[string]interface{}

type HTTPError struct {
	Message string `json:"RESPONSE CODE: 404"`
}

func TopSecretPost(ev events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var satelliteRequest topsecret.SatellitesRequest
	var reqBody HTTPError

	resp := events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Conten-Type":                  "aplication/json",
			"Access-control-Allow-Origin":  "*",
			"Access-control-Allow-Methods": "POST",
		},
	}

	err := json.Unmarshal([]byte(ev.Body), &satelliteRequest)
	if err != nil {
		resp.StatusCode = http.StatusBadRequest
		resp.Body = fmt.Sprintf(`{"": %s"}`, reqBody.Message)
		return resp, err
	}

	x, y, msg := topsecret.GetTopsecret(satelliteRequest)

	var resp1 topsecret.Position
	resp1.X = json.Number(fmt.Sprintf("%.1f", x))
	resp1.Y = json.Number(fmt.Sprintf("%.1f", y))

	resp.StatusCode = http.StatusOK
	resp.Body = (fmt.Sprintf(`RESPONSE CODE: %d\n , {"position": {%s} , "message":"%s" }`, resp.StatusCode, resp1, msg))

	return resp, nil
}
