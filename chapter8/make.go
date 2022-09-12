package main

import (
	"fmt"
	"math/rand"
)

func main() {
	slice := make([]int, 5, 10) //vai iniciar um slice com 5 posições com o tipo padrão do tipo escolhido e com capacidade de 10 elementos

	for index := range slice {
		slice[index] = rand.Intn(25)
	}

	fmt.Println(slice, len(slice), cap(slice))

	slice = append(slice, 20, 30, 40, 50, 70, 80) // quando passar o capacity ele vai criar um novo com a mesma capacidade ou seja +10 elementos, copiar os elementos do antigo e passar para o novo

	fmt.Println(slice, len(slice), cap(slice))

}
