package main

import "fmt"

func main() {

	var nota int
	players := map[string]int{
		"Santiago": 11,
	}

	fmt.Println("Informe sua nota: ")
	fmt.Scan(&nota)

	if nota >= 90 {
		fmt.Println("PARABÉNS! Você foi aprovado com distinção.")
	} else if nota >= 70 {
		fmt.Println("Aprovado!")
	} else {
		fmt.Println("Infelizmente você foi reprovado.")
	}
	fmt.Println()

	fmt.Println("Consultando players do map: ")
	//condição if em apenas uma linha.
	if value, ok := players["Santiago"]; ok {
		fmt.Println("Ponots:", value, ok)
	}
}
