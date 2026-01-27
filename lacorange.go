package main

import "fmt"

func main() {
	users := map[string]string{
		"nome":  "Santiago",
		"idade": "11",
	}

	for key, value := range users {
		fmt.Println(key, value)
	}
}
