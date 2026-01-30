package main

import "fmt"

func main() {

	var cresc [10]int

	fmt.Println("Informe os números que compõe sua lista: ")

	for i := 0; i < len(cresc); i++ {
		fmt.Print("Digite o ", i+1, "° número: ")
		fmt.Scan(&cresc[i])
	}
	fmt.Println("")
	fmt.Println("###################################")
	fmt.Println("Ordem inversa da lista informada: ")

	for i := 9; i >= 0; i-- {
		fmt.Println(cresc[i])
	}

}
