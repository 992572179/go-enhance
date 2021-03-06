package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	var a [10]int

	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				a[i]++
				//程序交出CPU
				runtime.Gosched()
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)

}
