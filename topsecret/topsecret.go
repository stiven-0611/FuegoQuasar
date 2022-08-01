package topsecret

import (
	"prueba/services"
	"strings"
)

func GetTopsecret(satellites SatellitesRequest) (x, y float32, message string) {

	var r1, r2, r3 float64
	var messages = make(map[string][]string)

	for _, satellite := range satellites.Satellites {

		switch strings.ToLower(satellite.Name) {
		case "kenobi":
			r1 = satellite.Distance
			messages["kenobi"] = satellite.Message
		case "skywalker":
			r2 = satellite.Distance
			messages["skywalker"] = satellite.Message
		case "sato":
			r3 = satellite.Distance
			messages["sato"] = satellite.Message
		}
	}

	resX, resY := services.GetLocation(r1, r2, r3)
	msg := services.GetMessage(messages["kenobi"], messages["skywalker"], messages["sato"])

	return resX, resY, msg
}
