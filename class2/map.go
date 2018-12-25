package main

import "fmt"

func main() {

	stu := map[string]string{
		"name":   "z3",
		"city":   "Los Angeles",
		"gender": "male",
	}
	fmt.Println(stu)

	delete(stu, "name")
	fmt.Println("after del:\t", stu)

	for k, v := range stu {
		fmt.Println(k, "\t", v)
	}

	dict := CreateMap()
	dict["k1"] = "v1"
	dict["k2"] = "v2"
}

func CreateMap() map[string]string {
	return make(map[string]string)
}
