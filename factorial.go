package main

import (
	"fmt"
	"strconv"
)

func fat(n int) int {
	if n == 0 {
		return 1
	}
	return n * fat(n-1)
}

func main() {

	var numberString string
	fmt.Println("########################################")
	fmt.Println("Vamos descobrir o Fatorial de um número")
	fmt.Println("Digite um número inteiro e positivo: ")
	fmt.Scan(&numberString)
	numberInt, erro := strconv.Atoi(numberString)

	for erro != nil || numberInt <= 0 {
		fmt.Println("\nErro!") // \n => pula linha
		fmt.Println("Digite um número inteiro e positivo: ")
		fmt.Scan(&numberInt)
		numberInt, erro = strconv.Atoi(numberString)
	}
	fmt.Println("\nFatorial de", numberInt, " é = ", fat(numberInt))
	fmt.Println("########################################")
}
