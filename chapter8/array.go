package main

import "fmt"

var x [5]int
var y [6]int

func main() {
	x[0] = 1
	x[1] = 2
	x[2] = 3
	fmt.Println(x)
	fmt.Println(x[0])
	fmt.Println(x[1])
	fmt.Println(x[2])
	fmt.Printf("%T\n", x)
	fmt.Println(len(x))
	fmt.Println(len(y))
}
