package main

import (
	"log"
	"math"
)

var coords = [][]float32{
	{-500, -200},
	{100, -100},
	{500, 100},
}

func main() {
	x, y := GetLocation(761.577310586391, 670.820393249937, 806.225774829855)

	log.Printf("Satellite coords: [%f, %f]", x, y)
}

func GetLocation(distances ...float32) (x, y float32) {
	dp2 := distances[0] * distances[0]
	dq2 := distances[1] * distances[1]
	dr2 := distances[2] * distances[2]

	px, py := coords[0][0], coords[0][1]
	qx, qy := coords[1][0], coords[1][1]
	rx, ry := coords[2][0], coords[2][1]

	px2, py2 := px*px, py*py
	qx2, qy2 := qx*qx, qy*qy

	rx2, ry2 := rx*rx, ry*ry

	x = (px2 + py2 - rx2 - ry2 + dr2 - dp2) * (py - qy)
	x = x - (px2+py2-qx2-qy2+dq2-dp2)*(py-ry)
	x = x / (2 * ((qx-px)*(py-ry) - (rx-px)*(py-qy)))

	y = x*2*(qx-px) + (px2 + py2 - qx2 - qy2 - dp2 + dq2)
	y = y / (2 * (py - qy))

	x, y = float32(math.Round(float64(x))), float32(math.Round(float64(y)))

	return
}
