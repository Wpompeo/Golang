package main

import "fmt"

type Cliente struct {
	Nome     string
	Idade    int
	Endereco Endereco
	Email    string
}

//Usar Estrutura dentro de struturas
type Endereco struct {
	Rua    string
	Numero int
	Cep    string
	Estado string
}

func main() {

	clienteOne := Cliente{
		Nome:  "João",
		Idade: 11,
	}

	clienteTwo := Cliente{
		Nome:  "Maria",
		Idade: 30,
	}

	//Criando cliente Three com Estruturas dentro de Estruturas
	clienteThree := Cliente{
		Nome:  "Pedro",
		Idade: 32,
		Endereco: Endereco{
			Rua:    "Rua das Flores",
			Numero: 123,
			Estado: "RS",
		},
	}

	fmt.Println(clienteOne)
	fmt.Println(clienteTwo)
	fmt.Println(clienteThree)

	//adicionar informacoes
	clienteOne.Email = "dev@gmail123.com"
	clienteTwo.Email = "devTwo@gmail1234.com"

	//Alterar dados estrutura
	clienteThree.Endereco.Numero = 1111

	fmt.Println(clienteOne.Email)
	fmt.Println(clienteTwo.Email)
	fmt.Println(clienteThree)

}
