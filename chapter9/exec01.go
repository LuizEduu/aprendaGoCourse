package main

import "fmt"

func main() {
	numbers := [5]int{1, 2, 3, 4, 5}

	for _, number := range numbers {
		fmt.Print(number)
	}

	fmt.Printf("\n%T\n", numbers)
}
