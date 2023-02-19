package main

import "fmt"

func main() {
	x := "texto\n" //literal string
	y := `texto2`  //raw (rune) string
	z := fmt.Sprintf(x, " ", y)
	w := fmt.Sprintf("%s %s", x, y)
	fmt.Println(z)
	fmt.Println(w)
}
