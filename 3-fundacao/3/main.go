package main

import "fmt"

// Essa é a sintaxe para criar um tipo
type customType string

var (
	message customType = "test"
)

func main() {
	fmt.Println(message)
}
