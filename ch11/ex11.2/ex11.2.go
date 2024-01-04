package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("숫자를 입력하세요(종료: 0): ")
		var num int
		_, err := fmt.Scanln(&num)
		if err != nil {
			fmt.Println("Error! 숫자 값을 입력해야합니다")
			_, _ = stdin.ReadString('\n')
			continue
		}
		fmt.Printf("입력한 숫자는 %d 입니다\n", num)
		if num == 0 {
			break
		}
	}
	fmt.Println("End")
}
