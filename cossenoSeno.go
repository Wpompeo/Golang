package main

import (
	"fmt"

	"math"
)

func main() {

	var angulo float64

	//entrada de dados
	fmt.Println("###########################")
	fmt.Println("Digite o ângulo desejado: ")
	fmt.Scan(&angulo)
	fmt.Println("###########################")

	//calculo e impressao do cosseno
	fmt.Println("Cosseno de", angulo, "é:", math.Cos(angulo))

	//calculo e impressao do seno
	fmt.Println("Seno de", angulo, "é:", math.Sin(angulo))

}
