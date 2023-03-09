package main

import (
	"fmt"
)

func main() {
	var variable int = 1
	fmt.Println(variable)

	variable = 2
	fmt.Println(variable)

	const constante int = 3
	fmt.Println(constante)

	//constante = 4   //!cannot assign to constante (constant 3 of type int)
	//fmt.Println(constante)
}
