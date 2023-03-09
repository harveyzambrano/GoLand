package main

import (
	"fmt"
	"reflect"
)

func main() {
	var cifra int = 99
	var palabra string = "Coto"

	fmt.Println(cifra)
	fmt.Println(palabra)
	fmt.Println(reflect.TypeOf((cifra)))
	fmt.Println(reflect.TypeOf((palabra)))
}
