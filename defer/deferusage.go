package main

import "fmt"

func main() {

	fmt.Println("a")

	//FILO
	defer fmt.Println("b")
	defer fmt.Println("c")

}
