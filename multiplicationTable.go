package main

import "fmt"

func main() {

	var number, result int

	//input for data
	fmt.Println("################SEJA BEM VINDO AO SIMULADOR DE TABUADA################")
	fmt.Println("Escolha um número positivo e inteiro: ")
	fmt.Scan(&number)

	//loop for repitition
	for i := 1; i <= 10; i++ {
		result = number * i
		fmt.Printf("%v X %v = %v\n", number, i, result)
	}
	fmt.Println("################PRONTO, ESSA É A TABUADA DO NÚMERO", number, "################")
}
