package main

import "fmt"

func main() {
	sabores := []string{"peperonni", "mozzarela", "quatro queijos", "calabresa", "frango com catupiry"}

	//	fatia := sabores[2:]

	/* fmt.Println(sabores[:])  //vai do inicio até o fim do slice
	fmt.Println(sabores[2:]) //vai da posição informada até o fim do slice
	fmt.Println(sabores[:2]) //vai do inicio do slice até o fim passando a posição informada */

	sabores = append(sabores[:2], sabores[3:]...) //remove um elemento do array
	fmt.Println(sabores)
}
