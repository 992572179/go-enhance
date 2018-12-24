package main

import "fmt"

//数组是值传递，方法中会拷贝一份，不会更改原值。
// [3]int和[5]int会被认为是不同的数组，不等价
func main() {

	//声明固定大小为3的整型数组
	ints := [3]int{5, 12, 13}
	//声明动态大小的整型数组
	arrays := [...]int{1, 3, 5, 7, 9}

	fmt.Println(ints, arrays)

	//数组的遍历
	/*for i := range arrays {
		fmt.Println(arrays[i])
	}*/

	/*for _, v := range arrays {
		fmt.Println(v)
	}*/

	IteArr(ints)
	//IteArr(arrays)
}

func IteArr(arr [3]int) {
	for i := range arr {
		fmt.Println(arr[i])
	}
}
