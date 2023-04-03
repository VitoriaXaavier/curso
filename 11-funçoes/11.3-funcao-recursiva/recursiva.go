package main

import (
	"fmt"
	
)
func Fibonacci(posicao uint) uint{
	if posicao <=1 {
		return posicao
	}
	return Fibonacci(posicao - 2) + Fibonacci(posicao - 1)
}

func Comma ( s string) string{
	n := len(s)
	if n <= 3{
		return s
	}
	return Comma(s[:n - 3]) + "," + s[n-3:]
}


func main() {

	posicao := uint(10)

	fmt.Println(Fibonacci(posicao))
	fmt.Println("-----------------")

	for i := uint(1); i <= posicao; i++ {
		fmt.Println(Fibonacci(i))
	}

	fmt.Println("-----------------")

	comma1 := Comma("vitoria")
	fmt.Println(comma1)

}
