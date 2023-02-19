package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	//https://www.youtube.com/watch?v=VkuZ4QMnoZM
	s := fmt.Sprintf("%v->%T\n%v->%T\n%v->%T\n", x, x, y, y, z, z) //Valor 0
	fmt.Println(s)
}
