package main

import "fmt"

func main() {
	var arr [3]int
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30

	lastItem := arr[len(arr)-1]

	fmt.Println(lastItem)

	for index, currentValue := range arr {
		fmt.Printf("valor no indice '%v': %v\n", index, currentValue)
	}
}
