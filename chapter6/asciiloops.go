package main

import "fmt"

func main() {
	for i := 33; i < 122; i++ {
		fmt.Printf("%v\n", string(rune(i))) //convert to rune um caracte em go que por baixo dos panos é um int32 e após isso converto para string
	}
}
