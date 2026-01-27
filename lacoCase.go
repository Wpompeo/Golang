package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Quando é Sábado? ")
	//declara var today a biblioteca time weekday
	today := time.Now().Weekday()

	switch time.Saturday {
	case today + 0:
		fmt.Println("É hoje!")
	case today + 1:
		fmt.Println("É amanhã!")
	case today + 2:
		fmt.Println("É em dois dias")
	default:
		fmt.Println("Está longe ainda!")
	}
}
