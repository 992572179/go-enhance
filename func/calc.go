package main

import "fmt"

func main() {
	i, e := calculate(3, 5, "+")
	if e != nil {
		panic(e)
	} else {
		fmt.Println(i)
	}
}

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
