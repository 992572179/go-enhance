package main

import "fmt"

func main() {

	stu := make(map[string]string)
	stu["name"] = "z3"
	stu["city"] = "Los Angeles"

	//map[name:z3 class:class1]
	fmt.Println(stu)

}
