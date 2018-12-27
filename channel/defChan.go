package main

import "fmt"

func main() {

	c := make(chan int)

	go func() {
		for {
			i := <-c
			fmt.Println(i)
		}
	}()
	//fmt.Printf("%T", c)

	//向chan写入数据
	c <- 1
	c <- 2

	buffedChannel := createBuffedChannel()
	closeChannel(buffedChannel)
	fmt.Println(buffedChannel)
}

//创建缓冲大小为3的channel
func createBuffedChannel() chan int {
	return make(chan int, 3)
}

func closeChannel(c chan int) {
	close(c)
}
