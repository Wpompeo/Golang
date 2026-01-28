package main

import "fmt"

func main() {

	var numberOne, numberTwo, result float32
	var operation string = " "

	//input data
	fmt.Println("###############BEM VINDO AO PROGRAMA DE OPERAÇÔES MATEMÁTICAS###############")
	fmt.Println("Digite o primeiro número de sua operação:")
	fmt.Scan(&numberOne)

	fmt.Println("Digite o segundo número de sua operação:")
	fmt.Scan(&numberTwo)

	fmt.Println("Informe o tipo de operação que deseja realizar. ( +, -, /, * ) ")
	fmt.Println(&operation)

	//validation conditional
	if operation != "+" && operation != "-" && operation != "*" && operation != "/" {
		fmt.Println("ATENÇÃO! Caracter inválido.")
		fmt.Println("Informe uma operação válida. ( +, -, /, * ) ")
		fmt.Scan(&operation)
	}

	//conditional and calculate
	if operation == "+" {
		result = numberOne + numberTwo
		fmt.Println("O resultado da soma dos números informados é:", result)
	} else if operation == "-" {
		result = numberOne - numberTwo
		fmt.Println("O resultado da sosubtração dos números informados é:", result)
	} else if operation == "*" {
		result = numberOne * numberTwo
		fmt.Println("O resultado da multiplicação dos números informados é:", result)
	} else if operation == "/" {
		result = numberOne / numberTwo
		fmt.Println("O resultado da divisão dos números informados é:", result)
	} else {
		fmt.Println("ATENÇÃO! Opção inválida.")

	}
	fmt.Println("###############FIM DE EXECUÇÃO DO PROGRAMA###############")

}
