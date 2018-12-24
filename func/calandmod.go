package main

import "fmt"

func main() {

	fmt.Println(CalAndMod(13, 5))

	fmt.Println(SumOfRanges(1, 2, 3, 4, 5))

}

//计算以及求余
func CalAndMod(a, b int) (int, int) {
	return a / b, a % b
}

//可变参数列表，但参数类型必须一致
func SumOfRanges(nums ...int) int {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	return sum
}
