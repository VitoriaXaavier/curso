package main

import "fmt"


func soma(n1 int, n2 int) int {
	return n1 + n2 
}

func calculoMatematico(n1,n2 int) (int,int){

	soma := n1 + n2
	subtracao := n1 - n2
	

	return soma, subtracao

}

func main() {

	soma := soma(10, 2)
	fmt.Println(soma)

	var f = func() {
		fmt.Println("f")
	}

	f()
	var s = func(txt string) string{
		return txt
	}
	resultado := s("texto")

	fmt.Println(resultado)

	resultadoSoma, _ := calculoMatematico(10,15)
	fmt.Println(resultadoSoma)


}