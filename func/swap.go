package main

import "fmt"

/**
GO语言默认是值传递
*/
func main() {

	a, b := 3, 5
	swap(a, b)
	fmt.Println(a, b)

	c, d := 12, 13
	swap2(&c, &d)
	fmt.Println(c, d)

	//0xc00004a0b8
	e := 100
	fmt.Println(&e)
}

//值传递
func swap(a, b int) {
	a, b = b, a
}

//引用传递
func swap2(a, b *int) {
	*a, *b = *b, *a
}
