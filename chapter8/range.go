package main

import "fmt"

func main() {
	slice := []string{"banana", "maçã", "jaca"}

	//slice[3] = "melancia" não funciona pois os slices tem tamanho fixo e para adicionar novos tem que utilizar append
	slice = append(slice, "melancia")

	for index, value := range slice { //executa o bloco de código para cada elemento do slice
		fmt.Println(index, value)
	}

	/* for index := range slice { utilizando apenas os index
		fmt.Println(index)
	} */

	/* for _, value := range slice { utilizando apenas os valores
		fmt.Println(value)
	} */

	slice2 := []int{20, 20, 20}

	total := 0
	for _, value := range slice2 {
		total += value
	}

	fmt.Println(total)

}
