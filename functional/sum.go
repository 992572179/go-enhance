package main

import "fmt"

func main() {
	value := sum()
	for i := 0; i < 6; i++ {
		fmt.Println(value(i))
	}
}

//函数式编程
func sum() func(int) int {
	result := 0
	return func(v int) int {
		result += v
		return result
	}
}
