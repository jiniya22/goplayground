package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("data.txt")
	if err != nil {
		fmt.Println("파일 생성 중 오류 발생")
		return
	}
	defer fmt.Println("defer 지연실행은 맨먼저 선언한것이 맨 나중에 실행됩니다.")
	defer file.Close()
	defer fmt.Println("맨 마지막에 사용한 defer 지연실행이지만 가장 먼저 실행됩니다.")

	fmt.Println("파일에 Hello jiniworld!를 썼습니다.")
	fmt.Fprintln(file, "Hello jiniworld!")
}
