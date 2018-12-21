package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	printFile("1.txt")
}

//打印目录中的某个文件的数据
func printFile(fileName string) {

	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
