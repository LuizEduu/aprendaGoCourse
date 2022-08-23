package main

import "fmt"

const (
	_ = iota
	b = iota
	c = iota
	x = iota
	y = iota
	z = iota
)

func main() {
	//iotas são valores sucessivos inteiros e não tipados, servem para constantes

	fmt.Println(b, c, x, y, z)

}
