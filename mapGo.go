package main

import "fmt"

func main() {
	var pessoas = map[string]int{}
	pessoas["Maria"] = 11
	pessoas["João"] = 30

	//Verifica se pessoa Joao existe no map
	if idade, ok := pessoas["João"]; ok {
		fmt.Println("Pessoa existente no map", idade, ok)

	} else {
		fmt.Println("Pessoa não existe no map")
	}
	//deleta Joao do map
	delete(pessoas, "João")
	//map atual
	fmt.Println(pessoas)

}
