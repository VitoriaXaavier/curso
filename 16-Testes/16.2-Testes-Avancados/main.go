package main

import (
	"teste-avancado/area"
)

func main() {

	r := area.Retangulo{10,15}
	area.EscreveArea(r)

	c := area.Circulo{10}
	area.EscreveArea(c)
}