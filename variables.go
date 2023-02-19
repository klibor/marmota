package main

import "fmt" // importa o fmt

var y = "Variável string" // Scope package (visivel em todo o código)

func main() {
	// := É chamado de marmota -> É o mascote do GoLang (gopher) ou o punisher / Só funciona dentro de code blocks
	p := 10.0
	x := 10
	fmt.Printf("x: %v, %T\n", x, x) // Exibe a var x com valor 10 e tipo int
	fmt.Printf("y: %v, %T\n", y, y) // Exibe em tela a var y com valor 'Variável string' e tipo string
	x, z := 20, 30
	fmt.Printf("x: %v, %T\n", x, x) // Exibe a var x com valor 20 e tipo int
	fmt.Printf("z: %v, %T\n", z, z) // Exibe a var z com valor 30 e tipo int
	fmt.Printf("p: %v, %T\n", p, p) // Exibe a var p com valor 10.0 e tipo float64
}
