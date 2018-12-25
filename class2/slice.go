package main

import "fmt"

func main() {

	arr := [...]int{3, 4, 5, 12, 13}
	fmt.Println(arr)

	s1 := arr[:]
	s2 := arr[0:3]
	s3 := s2[0:4]

	fmt.Println("s1:\t", s1)
	fmt.Println("s2:\t", s2)
	fmt.Println("s3:\t", s3)
	fmt.Printf("len: %d,cap: %d\n", len(s2), cap(s2))

	/**
	{3,4,5,12,13}
	-->[0:3]
	{3,4,5}
	len->3
	cap->5
	*/

	//切片是引用传递，直接更改原值
	ModifyArr(s2)
	fmt.Println("Modify s2\t", s2)
}

func ModifyArr(arr []int) {
	arr[0] = 0
}
