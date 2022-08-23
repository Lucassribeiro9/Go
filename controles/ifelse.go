package main

import "fmt"

func imprimirResultado(nota float64) {
	if nota >= 7 {
		fmt.Println("Aprovado")
	} else {
		fmt.Println("Reprovado")
	}

}

// else if

func notaParaConceito(n float64) string {
	if n >= 9 && n <= 10 {
		return "A"
	} else if n >= 8 && n <= 9 {
		return "B"
	} else if n >= 5 && 8 <= 10 {
		return "C"
	} else if n >= 3 && n <= 10 {
		return "D"
	} else {
		return "E"
	}
}

func main() {
	imprimirResultado(5.1)
	imprimirResultado(8)
	fmt.Println(notaParaConceito(6.1))
}
