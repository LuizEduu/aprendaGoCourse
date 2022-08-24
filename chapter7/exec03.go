package main

import "fmt"

func main() {
	yearNasc := 1998
	actualYear := 2022
	for yearNasc <= actualYear {
		fmt.Println(yearNasc)
		yearNasc++
	}
}
