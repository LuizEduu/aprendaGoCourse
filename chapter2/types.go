/*
	Go é uma linguagem de tipagem estática ou seja, quando algo é definido com um tipo, ela vai manter aquele tipo até o fim da execução do programa.
*/

package main

import "fmt"

var x int

//x = 20 //não é permitido pois uma var inicializada em escopo de package só pode ter valor atribuido dentro de um escopo de bloco
// primitive types int e suas variações, uint e suas variações, string, bool, float 2 variações, complex 2 variações

func main() {
	x = 20 //permitido
	fmt.Printf("%v, %T\n", x, x)
}
