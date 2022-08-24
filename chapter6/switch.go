package main

import "fmt"

func main() {
	quemtanoescritoriohoje := "joana"

	//switch é fallthrough false por padrão, caso habilitado ele vai executar o case e a próxima instrução sem verificar as condições para execução

	/* switch quemtanoescritoriohoje {
	case "zezinho":
		fmt.Println("quem tá no escritório é o Zezinho")
		fallthrough
	case "marquinhos":
		fmt.Println("quem tá no escritório é o Marquinhos")
	case "maria":
		fmt.Println("quem tá no escritório é a Maria")
	default:
		println("ninguém no escritório")
	} */

	switch quemtanoescritoriohoje {
	case "zezinho", "maria":
		fmt.Println("quem tá no escritório é o time 1")
	case "marquinhos", "joana":
		fmt.Println("quem tá no escritório é o time 2")
	default:
		println("ninguém no escritório")

	}
}
