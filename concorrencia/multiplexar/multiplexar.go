package main

import (
	"fmt"

	"github.com/Lucassribeiro9/html"
)

func encaminhar(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem
	}
}

// multiplexar

func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)
	return c
}

func main() {
	c := juntar(
		html.Titulo("https://www.cod3r.com.br", "https://www.espn.com.br"),
		html.Titulo("https://www.uol.com.br", "https://www.youtube.com"),
	)
	fmt.Println(<-c, "|", <-c)

}
