package main

import "fmt"

func main() {

	arr := [...]int{3, 4, 5, 12, 13}

	s1 := arr[0:3]

	fmt.Printf("s1: %d \t", s1)
	fmt.Println()

	splus := append(s1, 1001)

	fmt.Println(splus, cap(splus), len(splus))

}
