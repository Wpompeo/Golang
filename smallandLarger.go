package main

import "fmt"

func main() {

	var firstNumber, secondNumber int

	fmt.Println("#########################ATENÇÃO!#########################")

	fmt.Println("Você deve informar dois números, onde obrigatóriamente o segundo número deverá ser maior que o primeiro.")
	fmt.Println("Informe o primeiro número.")
	fmt.Scan(&firstNumber)
	fmt.Println("Informe o segundo número.")
	fmt.Scan(&secondNumber)

	for firstNumber > secondNumber {

		fmt.Println("Erro! O segundo número é menor que o primeiro.")
		fmt.Println("Digite novamente o segundo número.")
		fmt.Scan(&secondNumber)
	}
	fmt.Printf("O número %v é maior que o número %v.", secondNumber, firstNumber)
	fmt.Println()

	fmt.Println("#########################FIM DE COMPARAÇÃO!#########################")
}
