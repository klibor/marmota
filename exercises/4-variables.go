package main

import (
	"fmt"
)

type ex4Var int

var x ex4Var

func main() {
	//https://www.youtube.com/watch?v=5-0S-gefNe0
	fmt.Printf("%v -> %T\n", x, x)
	x = 42
	fmt.Printf("%v -> %T\n", x, x)
	y := ex4Var(x)
	fmt.Printf("%v -> %T\n", y, y)
}
