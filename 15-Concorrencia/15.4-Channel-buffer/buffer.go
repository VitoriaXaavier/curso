package main

import "fmt"

func main() {
	canal := make(chan string, 2) // cria um channel e especifica a quantidade de channels para não dar deadLock

	canal <- "Hello world"
	canal <- "Channel with buffer"

	mensagem := <- canal 
	mensagem2 := <- canal

	fmt.Println(mensagem,mensagem2)


}