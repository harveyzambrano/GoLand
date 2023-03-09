package main

import (
	"fmt"
	"reflect"
)

func main() {
	numeroEntero := 60
	numeroFlotante := 15.5
	var numero1 int8 = 127
	var numero2 int16 = 32767  // hasta el valor  2^16
	var numero3 uint16 = 10    // no incluye negativos
	var numero4 float32 = 15.5 // Para limitar el uso dememoria

	fmt.Println(numeroEntero)
	fmt.Println(reflect.TypeOf(numeroEntero))

	fmt.Println(numeroFlotante)
	fmt.Println(reflect.TypeOf(numeroFlotante))

	fmt.Println(numero1)
	fmt.Println(numero2)
	fmt.Println(numero3)
	fmt.Println(numero4)
}
