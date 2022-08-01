package topsecret

import (
	"encoding/json"
)

type Satellites struct {
	Name     string   `json:"name"`
	Distance float64  `json:"distance" `
	Message  []string `json:"message" `
}

type SatellitesRequest struct {
	Satellites []Satellites `json:"satellites"`
}

type Position struct {
	X json.Number `json:"x"`
	Y json.Number `json:"y"`
}

type SatellitesResponse struct {
	Position Position `json:"position"`
	Message  string   `json:"message"`
}
