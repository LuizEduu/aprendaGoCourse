package main

import "fmt"

func main() {
	x := 10

	if x < 10 {
		fmt.Println("entrou no if")
	} else if x > 10 {
		fmt.Println("caiu no if else")
	} else {
		fmt.Println("caiu no else")
	}

}
