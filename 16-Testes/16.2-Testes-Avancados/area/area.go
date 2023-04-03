package area

import (
	"fmt"
	"math"
)

type Retangulo struct{
	Altura float64
	Largura float64
}
type Circulo struct {
	Raio float64
}

type Forma interface {
	Area() float64
}

func (r Retangulo) Area() float64 {
	return r.Altura * r.Largura
}
func ( c Circulo) Area() float64 {
	return math.Pi * math.Pow(c.Raio,2)
}

func EscreveArea (f Forma) {
	fmt.Printf("A área é de %0.2f \n", f.Area())
}
