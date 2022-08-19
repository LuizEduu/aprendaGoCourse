package main

import "fmt"

var x int     //declaração
var y string  //declaração
var z bool    //declaração
var a float64 //declaração

func main() {
	//se á variável é declarada e não é atribuido nenhum valor para ela, o padrão será o 0 (que varia dependendo do tipo da variável)
	//valor zero de int = 0
	//valor zero de string = "" string vazia
	//valor zero de string = 0.0
	//valor zero de bool = false
	// pointers, functions, interfaces, slices, channels, maps: nil (nada)

	fmt.Printf("%v, %T\n", x, x)
	fmt.Printf("%v, %T\n", y, y)
	fmt.Printf("%v, %T\n", z, z)
	fmt.Printf("%v, %T\n", a, a)
}
