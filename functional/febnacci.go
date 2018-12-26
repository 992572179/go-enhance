package main

import "fmt"

func main() {

	f := feb()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

}

func feb() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
