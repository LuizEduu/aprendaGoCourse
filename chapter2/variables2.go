package main

import "fmt"

func main() {
	x := 10 // Short variable declaration operator, serve para declarar uma variável ele declara uma variável, atribui o nome a ela e o tipo
	//ele tem tipagem automática, ou seja aribui o tipo pelo valor que foi dado aquela variável

	y := "fala tu, GO!"
	z := true
	w := 50.55

	fmt.Printf("x: %v, %T\n", x, x) //%valor da variável, %T tipo da variável
	fmt.Printf("y: %v, %T\n", y, y)
	fmt.Printf("z: %v, %T\n", z, z)
	fmt.Printf("w: %v, %T\n", w, w)

	x, a := 40, 255

	fmt.Println(x, a)
}
