package main

type customType string // Dessa forma é possível criar um tipo customizado, nesse exemplo está sendo criado o tipo "customType" que recebe uma string

var a customType = "teste"

func main() {
	println(a)
}
