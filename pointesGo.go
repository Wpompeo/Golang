package main

import "fmt"

type Pessoa struct {
	Nome string
}

func main() {
	var p1 Pessoa = Pessoa{Nome: "Maria"}
	var p2 Pessoa = Pessoa{Nome: "João"}
	fmt.Println("Nome p1 inicial: ", p1)

	var p3 *Pessoa = &p1
	p3.Nome = "Joana"
	fmt.Println("Nome p1 alterado via *p3: ", p1)
	fmt.Println(p2)
	fmt.Println(p3)

}
