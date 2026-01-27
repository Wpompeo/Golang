package main

import "fmt"

func main() {

	resultSum := Soma(2, 7)
	resultDiv := div(10, 2)

	fmt.Println("Resultado da soma é: ", resultSum)
	fmt.Println("Resultado da soma é: ", resultDiv)

}

/*Precisa ser declarado o tipo de parametro de entrada
funcao publica, pois é inicializada com letra Maiscula.*/
func Soma(a, b int) int {
	return a + b
}

//funcao privada, inicializada com letra miniscula.
func div(a, b int) int {
	return a / b
}
