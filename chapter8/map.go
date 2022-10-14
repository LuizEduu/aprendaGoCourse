package main

import "fmt"

func main() {
	friends := map[string]string{
		"João":    "1345562221",
		"armando": "4567832",
	}

	fmt.Println(friends)

	friends["Marcos"] = "214151"

	// comma ok idiom => return true with value exists or false with not exists
	if exists, ok := friends["Marcos"]; !ok {
		fmt.Println("não existe, Po!")
	} else {
		fmt.Println(exists)
	}

}
