package main

import "fmt"

type Target struct {
	name string
}

//方法属于结构体Target，只有结构体实例才能调用
func (target *Target) fooFunc() {
	fmt.Println("Target's target()...")
}

func (target *Target) getName() string {
	return target.name
}

func getName2(tg Target) string {
	return tg.name
}

func others() {
	fmt.Println("others")
}

func main() {
	others()

	//通过结构体实例调用方法
	target := Target{}
	target.fooFunc()

	tg2 := Target{
		name: "z3",
	}

	tg3 := Target{
		name: "List",
	}
	fmt.Println(tg2.getName())
	fmt.Println(getName2(tg3))

}
