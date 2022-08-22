package main

import "fmt"

func main() {
	s := "Hello, áêó"
	// %i, %T
	// Raw string literals
	// conversion to slice of bytes: []byte(x)
	// %#U, %#x

	for _, v := range s {
		fmt.Printf("%b / %T / %#U / %#x\n", v, v, v, v)
	} //com range ele vai ler caractere por caractere utilizando utf8

	fmt.Println("")

	for i := 0; i < len(s); i++ {
		fmt.Printf("%b - %T - %#U - %#x\n", s[i], s[i], s[i], s[i])
	} // com for loop ele vai ler byte por byte e pode dar problemas em caracteres com acento

}
