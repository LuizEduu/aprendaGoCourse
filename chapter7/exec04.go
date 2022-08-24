package main

import "fmt"

func main() {
	yearNasc := 1998
	actualYear := 2022
	for {

		if yearNasc > actualYear {
			break
		}
		fmt.Println(yearNasc)
		yearNasc++
	}
}
