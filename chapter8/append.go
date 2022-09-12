package main

import "fmt"

func main() {
	oneSlice := []string{"a", "b", "c"}
	twoSlice := []string{"d", "e", "f"}

	appendSlice := append(oneSlice, twoSlice...)

	fmt.Println(appendSlice)
}
