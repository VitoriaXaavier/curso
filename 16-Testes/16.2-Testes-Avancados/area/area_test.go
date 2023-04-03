package area

import (
	"testing"	
	"math"
)
func TestArea(t *testing.T) {
	t.Run("Retangulo", func(t *testing.T){
		ret := Retangulo{10,12}
		areaEsperada := float64(120)
		areaRecebida := ret.Area()

		if areaEsperada != areaRecebida {
			t.Errorf("A area esperada era de %f, mas foi recebida a area de %f", areaEsperada,areaRecebida)
		}
	})


	t.Run("Circulo", func(t *testing.T){
		cir:= Circulo{10}
		areaEsperada := float64(math.Pi * 100) // pi * o raio ao quadrado
		areaRecebida := cir.Area()

		if areaEsperada != areaRecebida{
			t.Errorf("A area esperada era de %0.2f, mas foi recebida a area de %0.2f", areaEsperada,areaRecebida)
		}
	})
}