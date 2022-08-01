package services

import (
	"fmt"
	"math"
	"strconv"
)

type Point struct {
	X float64
	Y float64
}

var (
	Kenobi    = Point{-500, -200}
	Skywalker = Point{100, -100}
	Sato      = Point{500, 100}
)

func GetLocation(distances ...float64) (x, y float32) {

	if distances == nil {
		return 0.0, 0.0
	}

	if len(distances) < 3 {
		return 0.0, 0.0
	}

	var (
		x1 = Kenobi.X
		y1 = Kenobi.Y

		x2 = Skywalker.X
		y2 = Skywalker.Y

		x3 = Sato.X
		y3 = Sato.Y

		r1 = distances[0]
		r2 = distances[1]
		r3 = distances[2]
	)

	var expx1 = (((math.Pow(r1, 2)-math.Pow(r2, 2))+(math.Pow(x2, 2)-math.Pow(x1, 2))+(math.Pow(y2, 2)-math.Pow(y1, 2)))*((2*y3)-(2*y2)) - ((math.Pow(r2, 2)-math.Pow(r3, 2))+(math.Pow(x3, 2)-math.Pow(x2, 2))+(math.Pow(y3, 2)-math.Pow(y2, 2)))*((2*y2)-(2*y1)))
	var expx2 = (((2*x2)-(2*x3))*((2*y2)-(2*y1)) - ((2*x1)-(2*x2))*((2*y3)-(2*y2)))
	var puntoX = expx1 / expx2
	var expy1 = ((math.Pow(r1, 2) - math.Pow(r2, 2)) + (math.Pow(x2, 2) - math.Pow(x1, 2)) + (math.Pow(y2, 2) - math.Pow(y1, 2)) + puntoX*((2*x1)-(2*x2)))
	var expy2 = ((2 * y2) - (2 * y1))
	var puntoY = expy1 / expy2
	var resXF string = fmt.Sprintf("%.1f", puntoX)
	var resYF string = fmt.Sprintf("%.1f", puntoY)

	var resY64 float64
	var resX64 float64

	var err error

	resY64, err = strconv.ParseFloat(resYF, 32)
	if err != nil {
		return 0.0, 0.0
	}
	resX64, err = strconv.ParseFloat(resXF, 32)
	if err != nil {
		return 0.0, 0.0
	}

	return float32(resX64), float32(resY64)
}
