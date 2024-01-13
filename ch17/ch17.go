package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var stdin = bufio.NewReader(os.Stdin)

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	cnt := 1
	r := rand.Intn(100)
	for {
		fmt.Print("숫자값을 입력하세요:")
		n, err := InputNum()
		if err != nil {
			fmt.Println("숫자값을 입력해주세요.")
			continue
		}
		if n == r {
			fmt.Println("숫자를 맞췄습니다. 축하합니다. 시도한 횟수:", cnt)
			break
		} else if n > r {
			fmt.Println("입력하신 숫자가 더 큽니다.")
		} else {
			fmt.Println("입력하신 숫자가 더 작습니다.")
		}
		cnt++
	}
}

func InputNum() (int, error) {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		stdin.ReadString('\n')
	}
	return n, err
}
