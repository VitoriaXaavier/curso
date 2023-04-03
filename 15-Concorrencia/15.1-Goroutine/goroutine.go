package main

import (
	"fmt"
	"time"
)

func main () {

	go escreve("Helo world!")
	escreve("Goroutine em Golang")
}


func escreve ( texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}