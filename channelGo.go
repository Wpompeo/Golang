package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 3)

	go func() {
		//ler no channel
		ch <- 10
		ch <- 12
		ch <- 13
	}()

	//ler valor de uma channel
	time.Sleep(time.Second * 1)
	valor := <-ch
	fmt.Println("Valor do canal: ", valor)

}
