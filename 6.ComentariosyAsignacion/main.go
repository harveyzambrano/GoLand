package main

import "fmt"

//forma No.1
var num int

//forma No.2
var texto string = "Harvey"

func main() {
	fmt.Println(num) //pprque no le habiamos dado un valor
	fmt.Println(texto)

	//forma No.3
	var texto2 string
	fmt.Println(texto2)

	//forma No.4
	var num2 int
	fmt.Println(num2)

	//forma No.5
	variable := 'h'
	fmt.Println("la h: ", variable)

	variable = 'a'
	fmt.Println("la a: ", variable)

	//forma No.6
	var v1, v2, v3 int
	fmt.Println("la v1: ", v1)
	fmt.Println("la v2: ", v2)
	fmt.Println("la v3: ", v3)

	//forma No.7
	variable1, variable2, variable3 := 10, 20, 30
	fmt.Println(variable1)
	fmt.Println(variable2)
	fmt.Println(variable3)
}
