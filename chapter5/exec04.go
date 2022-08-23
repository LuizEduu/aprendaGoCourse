package main

import "fmt"

func main() {
	a := 100

	fmt.Printf("%d\t%b\t%#x\n", a, a, a)
	b := a << 1 //deslocar um bit a esquerda multiplica por 2
	fmt.Printf("%d\t%b\t%#x\n", b, b, b)
}
