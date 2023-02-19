package main

import "fmt" // importa o fmt

var y = "Variável string" // Package scope (visivel em todo o código)
var q string              // Definindo variavel com um tipo associado

// Valor 0 de variaveis sem inicialização explicita
var value0_bool bool
var value0_int int
var value0_float float32
var value0_string string

func main() {
	// := É chamado de marmota -> É o mascote do GoLang (gopher) ou o punisher / Só funciona dentro de code blocks
	n := true
	o := []string{"um", "dois"}
	p := 10.0
	q = "Nova variável string" // Adicionando valor a variável q definida no package scope
	x := 10
	fmt.Printf("x: %v, %T\n", x, x) // Exibe a var x com valor 10 e tipo int
	fmt.Printf("y: %v, %T\n", y, y) // Exibe em tela a var y com valor 'Variável string' e tipo string
	x, z := 20, 30
	fmt.Printf("x: %v, %T\n", x, x) // Exibe a var x com valor 20 e tipo int
	fmt.Printf("z: %v, %T\n", z, z) // Exibe a var z com valor 30 e tipo int
	fmt.Printf("p: %v, %T\n", p, p) // Exibe a var p com valor 10.0 e tipo float64
	fmt.Printf("n: %v, %T\n", n, n) // Exibe a var p com valor true e tipo bool
	fmt.Printf("o: %v, %T\n", o, o) // Exibe a var o com valor ['um','dois'] e tipo array de strings

	// Imprime value0 vars
	fmt.Printf("value0_bool: %v, %T\n", value0_bool, value0_bool)       // Exibe a var value0_bool com valor false e tipo bool
	fmt.Printf("value0_int: %v, %T\n", value0_int, value0_int)          // Exibe a var value0_int com valor 0 e tipo int
	fmt.Printf("value0_float: %v, %T\n", value0_float, value0_float)    // Exibe a var value0_float com valor 0 e tipo float32
	fmt.Printf("value0_string: %v, %T\n", value0_string, value0_string) // Exibe a var value0_float com valor "" e tipo string
}
