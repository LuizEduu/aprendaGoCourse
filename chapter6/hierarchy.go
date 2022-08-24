package main

import "fmt"

func main() {
	for hours := 1; hours <= 24; hours++ {
		fmt.Println("Hora:", hours)
		for minutes := 1; minutes <= 60; minutes++ {
			fmt.Print(" ", minutes)
		}
		fmt.Println()
	}

}
