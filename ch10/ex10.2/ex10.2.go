package main

import "fmt"

func main() {
	var temp int
	fmt.Scanf("%d", &temp)

	switch {
	case temp < 10 || temp > 30:
		fmt.Println("놀러가기 좋지 않은 날씨입니다")
	case temp > 15 && temp < 20:
		fmt.Println("적당한 겉옷을 입고 나가면 좋을 날씨입니다")
	default:
		fmt.Println("기타")
	}
}
