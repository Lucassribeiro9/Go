package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5}
	for i, numero := range numeros { // retorna o indice e o array que está percorrendo
		fmt.Printf("%d) %d\n", i+1, numero)
	}

	for _, num := range numeros {
		fmt.Println(num)
	}
}
