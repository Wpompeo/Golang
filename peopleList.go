package main

import (
	"fmt"
	"strings"
)

func main() {

	const quantityVector int = 3
	var name [quantityVector]string
	var age [quantityVector]int

	for i := 0; i < quantityVector; i++ {

		fmt.Printf("Informe o %v° nome: ", i+1)
		fmt.Scan(&name[i])

		fmt.Printf("Informe a idade de %v: ", name[i])
		fmt.Scan(&age[i])

		fmt.Println()
	}

	var searchs []string

	for consultation := " "; consultation != "N"; {

		var search string = " "

		consultation = " "

		fmt.Println("\n--CONSULTAS--")

		for search == " " {

			fmt.Println("Nome a ser consultado: ")

			fmt.Scan(&search)

			for _, nameItem := range searchs {

				if search == nameItem {

					fmt.Printf("\nNome já pesquisado anteriormente!\n")

					search = " "

					break
				}
			}
		}

		searchs = append(searchs, search)

		var occurrence int = 0

		for i, nameItem := range name {

			if search == nameItem {

				fmt.Printf("A pessoa \"%v\" foi encontrada e possui %v anos.\n,", nameItem, age[i])

				occurrence++
			}
		}

		if occurrence == 0 {

			fmt.Println("\nPessoa não localizada.")
		} else {
			fmt.Printf("\nForam encontradas %v ocorrências do nome informado.\n", occurrence)
		}

		for search != "S" && search != "N" {

			fmt.Println("\nDeseja realizar uma nova consulta? S/N")

			fmt.Scan(&search)

			search = strings.ToUpper(search)
		}

	}

}
