package main

import (
	"fmt"
)

type ex4Var int

var x ex4Var
var y int

func main() {
	//https://www.youtube.com/watch?v=O0r318FN_Uw
	fmt.Printf("Valor de x = %v -> Tipo de x = %T\n", x, x)
	x = 42
	fmt.Printf("Valor de x = %v -> Tipo de x = %T\n", x, x)
	y = int(x) //Conversão
	fmt.Printf("Valor de y = %v -> Tipo de y = %T\n", y, y)
}
