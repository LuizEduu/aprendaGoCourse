package main

import (
	"fmt"
	"runtime"
)

func main() {
	/*
		Integers:
		int & uint int8,int16,int32,int64 -> seguem o padrão da word da arquitetura do processador, ou seja se for 32bits será um int de 32 se for 64 será 64 bits
		todos que vem com uint são unsigned ou seja o u na frente indica que é numeros sem sinais.
		rune -> int32(UTF8)
		cada caractere no go é chamado de rune e cada rune ocupa o espaço de um int32
		float -> float32 ou float64, novamente segue o padrão da arquitetura do processador, se for 32, float32 64 float64
	*/

	a := "e"
	b := "é"

	fmt.Printf("%v,%v\n", a, b)

	c := []byte(a)
	d := []byte(b)
	fmt.Printf("%v,%v\n", c, d)

	e := 10
	f := 10.0

	fmt.Printf("%v, %T\n%v,%T \n", e, e, f, f)

	//h := []byte(e) não é possível converter tipos numericos em array de bites

	fmt.Println(runtime.GOOS)   //view PC OS
	fmt.Println(runtime.GOARCH) //view PC archtecture
}
