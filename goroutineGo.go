package main

import (
	"fmt"
	"time"
)

func exibirMsg() {
	fmt.Println("Olá de uma gorouitne!")
}

func main() {
	go exibirMsg()
	//tempo necessario para nao finalizar main antes de executar a goRoutine
	time.Sleep(1 * time.Second)
	fmt.Println("Olá main function..")

}
