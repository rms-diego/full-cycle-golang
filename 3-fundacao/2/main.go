package main

import "fmt"

// pode declarar multiplas variaveis ou variveis sozinhas, os tipos sao inferidos ou declarados

// multiplas
var (
	boolean bool    = true
	integer int     = 20
	decimal float64 = 3.22222222222
	text    string  = "hello, world!"
)

// sozinha
var test string = "message 1"

const test2 = "message 2"

func main() {
	fmt.Println(boolean)
	fmt.Println(integer)
	fmt.Println(decimal)
	fmt.Println(text)

	fmt.Println(test)
	fmt.Println(test2)
}
