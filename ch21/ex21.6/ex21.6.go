package main

import (
	"fmt"
	"os"
)

type Writer func(string)

func writeHello(writer Writer) {
	writer("Hello")
}

func main() {
	file, err := os.Create("data.txt")
	if err != nil {
		fmt.Println("파일 생성 실패")
		return
	}
	defer file.Close()
	writeHello(func(msg string) {
		fmt.Fprintln(file, msg)
	})
}
