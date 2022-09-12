package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)
	slice = append(slice, 6)
	fmt.Println(slice)

	//slice[20] = 20 não funciona pois por baixo dos panos slice utiliza arrays e o array subjacente ou seja o anterior não tinha dados até a posição 19
}
