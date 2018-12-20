package main

import "fmt"

func addSum() int {
	sum := 0
	for i := 0; i <= 100; i++ {
		sum += i
	}
	return sum
}

//死循环
func loops() {
	for {
		fmt.Println("hello world")
	}
}

func main() {

	fmt.Println(addSum())

	//loops()

}
