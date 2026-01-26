package main

import "fmt"

func main() {
	var gavetas [2]string
	var portas []string
	gavetas[0] = "Copos"
	gavetas[1] = "Talheres"
	fmt.Println("O que temos nas gavetas? " + gavetas[0] + " e " + gavetas[1])

	portas = append(portas, "Pratos", "Toalhas", "Colheres")
	fmt.Println("Tamanho do meu Slice é:", len(portas))
	fmt.Println("O que temos nas portas do armário?", portas)
	fmt.Println("Divisão do Slice, ponto inicial e ponto final:", portas[1:3])
	//removendo itens e criando um novo Slice
	portas = portas[:2]
	fmt.Println("Novo Slice é:", portas)

}
