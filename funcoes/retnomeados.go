package main

import "fmt"

func trocar(p1, p2 int) (segundo, primeiro int) {
	primeiro = p1
	segundo = p2
	return
}

func main() {
	x, y := trocar(5, 4)
	fmt.Println(x, y)
	segundo, primeiro := trocar(7, 1)
	fmt.Println(segundo, primeiro)
}
