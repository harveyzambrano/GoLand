package main

import (
	"fmt"
)

func main() {
	var num1 int = 20
	var num2 int = 10

	fmt.Println(num1 + num2)
	fmt.Println(num1 - num2)
	fmt.Println(num1 * num2)
	fmt.Println(num1 / num2)
	fmt.Println(num1 % num2)

	num2++
	fmt.Println(num2)

	num1--
	fmt.Println(num1)

	fmt.Println("Hola ", num1)
}
