package main

import "fmt"

func main() {
	i := add(3, 5)
	fmt.Println(i)
}

func add(a, b int) int {
	return a + b
}
