package main

import (
	"fmt"
	"math"
)

func main() {
	if i, err := calculate(3, 5, "-"); err != nil {
		panic(err)
	} else {
		fmt.Println("result:", i)
	}

	//把函数名作为参数
	fmt.Println(calculation(pow, 3, 4))

}

//对两个数的计算操作
func calculate(a, b int, operate string) (int, error) {
	switch operate {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf("%s", "Operation not support!"+"\t"+operate)
	}
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

//函数式编程
func calculation(op func(int, int) int, a, b int) int {
	return op(a, b)
}
