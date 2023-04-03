package main

import "fmt"

func main() {
	var var1 string = "variavel string sintaxe longa"
	var2 := "varivel string sintaxe curta"
	
	fmt.Println(var1)
	fmt.Println(var2)

	var (
		var3 string = "variavel string multipla atribuição "
		var4 string = "variavel string multipla atribuição"
	)

	fmt.Println(var3, var4)

	var5, var6 := "variavel string inferencia de tipo, atribuição multipla", "variavel string inferencia de tipo, atribuição multipla"

	fmt.Println(var5,var6)


	//troca o valor das variaveis
	var1, var3 = var3, var1
	fmt.Println(var1,var3)
}