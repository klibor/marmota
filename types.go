package main

import "fmt"

//https://go.dev/ref/spec#Types
type hotdog int

var b hotdog = 10

func main() {
	x := 11
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", x)
	//x = b -> Erro, apesar de ser derivada de int, hotdog é um tipo diferente
	x = int(b) // Conversão de tipos -> type(value)
	fmt.Printf("%T", x)
	b = hotdog(x)
}
