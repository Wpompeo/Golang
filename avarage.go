package main

import "fmt"

//funcao calcula media
func Media(total float32) float32 {
	return total / 4
}

func main() {

	var n, total float32
	bim := 1

	// solicitação de notas ao usuário
	for bim >= 1 && bim <= 4 {
		fmt.Printf("Qual a nota do %dº Bimestre (0 a 10): ", bim)
		fmt.Scan(&n)

		// validação da nota
		for n < 0 || n > 10 {
			fmt.Print("Nota inválida! Digite um valor entre 0 e 10: ")
			fmt.Scan(&n)
		}

		total += n
		bim++
	}

	// chamando e apresentando a função média
	result := Media(total)

	if result >= 7 {
		fmt.Println("\nParabéns, você foi aprovado!")
		fmt.Print("\nSua média final é: ")
		fmt.Println(result)
	} else {
		fmt.Print("\nInfelizmente você não atingiu a média minima para a provação (7): ")
		fmt.Println(result)
	}

}
