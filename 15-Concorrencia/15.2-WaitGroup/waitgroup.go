package main

import (
	"fmt"
	"sync"
	"time"
)

func main () {

	var waitGroup sync.WaitGroup
	waitGroup.Add(2) // diz quantas goroutines tem para executar

	// executa concorrentemente as funções 
	go func () {
		escreve("Hello world")
		waitGroup.Done() // -1 no add
	}()

	go func () {
		escreve("Goroutine com WaitGroup")
		waitGroup.Done() // -1 no add
	}()

	waitGroup.Wait() // Acaba o programa quando o add for == a 0
}


func escreve ( texto string) {
	for i := 0; i <=5; i++{
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}