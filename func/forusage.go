package main

import (
	"fmt"
	"strconv"
)

//for循环的高阶操作：没有初始条件
func main() {
	s := iToBinary(8102)
	fmt.Println(s)
}

//整型转二进制
func iToBinary(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		temp := n % 2
		result = strconv.Itoa(temp) + result
	}
	return result
}
