package main

import "fmt"

func main() {
	i := 1

	// Go não possui aritmetica de ponteiros
	var p *int = nil

	p = &i // pegar o endereço da variável
	*p++   // desreferenciando o ponteiro para pegar o valor

	fmt.Println(p, *p, i, &i)

}
