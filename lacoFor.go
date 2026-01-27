package main

import "fmt"

func main() {

	sum := 0
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		sum += i
	}
	fmt.Println(sum)
	fmt.Println()

	sumTwo := 0
	//while nao está presente em GO, essa declaraçao representa a logica while
	for sumTwo < 20 {
		fmt.Println("Loop while no for")
		sumTwo += 2
	}
	fmt.Println(sumTwo)
}
