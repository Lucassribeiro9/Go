package main

import "fmt"

func inc1(n int) {
	n++
}

// ponteiro é representado por *

func inc2(n *int) {
	// * é usado para desreferenciar, ou seja, ter acesso ao valor no qual o ponteiro aponta
	*n++
}
func main() {
	n := 1
	inc1(n)
	fmt.Println(n)
	// & usado para obter o endereço da variável
	inc2(&n) // por referencia
	fmt.Println(n)
}
