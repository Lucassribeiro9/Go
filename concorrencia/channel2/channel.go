package main

import (
	"fmt"
	"time"
)

func doisTresQuatroVzs(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(3 * time.Second)
	c <- 4 * base
}
func main() {
	ch := make(chan int)
	go doisTresQuatroVzs(2, ch)

	a, b := <-ch, <-ch
	fmt.Println(a, b)

	fmt.Println(<-ch)
}
