package main

import "fmt"

func main() {
	s := make([]int, 10) // tipo do slice e o número de elementos
	s[9] = 12
	fmt.Println(s)
}
