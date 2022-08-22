package main

import "fmt"

var x bool //tipo booleano da linguagem

func main() {
	fmt.Println(x) //zero value de bool é false
	x = 10 == 10
	fmt.Println(x)
}
