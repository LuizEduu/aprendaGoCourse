package main

import "fmt"

const (
	_ = 2022 + iota //2022 + 0
	a
	b
	c
	d
)

func main() {

	fmt.Println(a, b, c, d)

	fmt.Printf("%#x\n", 31337)
}
