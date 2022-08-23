package main

import "fmt"

/* const (
	_  = iota             // 0
	KB = 1 << (iota * 10) // 1 << (1 * 10)
	MB = 1 << (iota * 10) // 1 << (2 * 10)
	GB = 1 << (iota * 10) // 1 << (3 * 10)
	TB = 1 << (iota * 10) // 1 << (4 * 10)
) */

const (
	_  = iota             // ignorando o 0
	KB = 1 << (iota * 10) // 1 << (1 * 10) //iota = 1
	MB                    //iota = 2
	GB                    //iota = 3
	TB                    //iota = 4
)

func main() {
	/*x := 24
	y := x << 2 //left bitwise operation, deslocamento de um bit de x para a esquerda
	z := x >> 2 //right bitwise operation, deslocamento de um bit de x para a direita
	fmt.Printf("%b\t%v\n", x, x)
	fmt.Printf("%b\t%v\n", y, y)
	fmt.Printf("%b\t%v\n", z, z)*/

	fmt.Println("binary\t\t\t\t\tdecimal")
	fmt.Printf("%b\t\t\t\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t\t\t", MB)
	fmt.Printf("%d\n", MB)
	fmt.Printf("%b\t\t", GB)
	fmt.Printf("%d\n", GB)
	fmt.Printf("%b\t", TB)
	fmt.Printf("%d\n", TB)
}
