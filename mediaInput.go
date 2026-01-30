package main

import "fmt"

func main() {

	var numbers [11]int
	var cont, sum, media int

	fmt.Println("Bem vindo ao programa de soma e media: ")

	for cont = 1; cont <= 10; cont++ {
		fmt.Printf("Digite o %v número: ", cont)
		fmt.Scan(&numbers[cont])

	}

	//laco para a soma e média dos 10 números

	for cont := 0; cont <= 10; cont++ {
		sum = sum + numbers[cont]

		//calcula media

		media = sum / 10
	}

	fmt.Printf("A soma dos valores informado é: %v\n", sum)
	fmt.Printf("A média dos valores informado é: %v ", media)

}
