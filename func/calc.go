package main

import (
	"fmt"
)

func main() {
	if i, err := calculate(3, 5, "-"); err != nil {
		panic(err)
	} else {
		fmt.Println("result:", i)
	}
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

func calcbyfunc(op func(int, int) int, a, b int) int {
	return op(a, b)
}
