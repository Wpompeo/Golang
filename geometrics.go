package main

import "fmt"

func main() {

	//declaracao variaveis

	var altura, base, area float32

	//entrada de dados

	fmt.Print("Valor da base do triângulo: ")

	fmt.Scan(&base)

	fmt.Print("Valor da altura do triângulo: ")

	fmt.Scan(&altura)

	//calculo da area

	area = (base * altura) / 2

	//saida

	fmt.Printf(" A area do triâgulo é: %v", area)

}
