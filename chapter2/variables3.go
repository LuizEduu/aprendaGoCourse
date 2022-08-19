package main

import "fmt"

var y = 10 //package scope

func main() {
	print(y)
}

func print(x int) {
	fmt.Println(x)
	fmt.Println(y)
}
