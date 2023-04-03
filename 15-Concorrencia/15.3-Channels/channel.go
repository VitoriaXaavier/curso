package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)

	go escreve("Hello world", canal)
	go escreve("Goroutine channels", canal)

	// for {
	// 	mensagem, aberto  := <-canal
	// 	if !aberto {
	// 		break
	// 	}
	// 	fmt.Println(mensagem)

	// }

	for mensagem := range canal {
		fmt.Println(mensagem)
	}

}

func escreve ( texto string, canal chan string) {
	for i := 0; i < 5; i++{
		canal <- texto
		time.Sleep(time.Second)
	}
	close(canal)
}