package main

import "fmt"

func getAge() int {
	return 30
}
func main() {
	switch age := getAge(); age {
	case 18:
		fmt.Println("18살!")
	case 30:
		fmt.Println("30!!!")
	default:
		fmt.Println("내 나이는", age)
	}

	// switch 블록 밖에서는 age 변수 사용 불가
}
