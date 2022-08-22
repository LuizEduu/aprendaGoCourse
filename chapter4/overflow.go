package main

import "fmt"

var i uint16 = 65535

func main() {
	i++
	fmt.Printf("%v, %T\n", i, i) //ao inves de dar overflow coloca o valor para 0
}
