package main

import "fmt"

type Data struct {
	value string
	score [4]int
}

func main() {
	var data Data
	fmt.Printf("data.value의 주소값: %p\n", &data.value)
	ChangeData(data)
	fmt.Printf("data의 value: %s, score[3]: %d\n", data.value, data.score[3])
}

func ChangeData(data Data) {
	fmt.Printf("ChangeData 내의 data.value의 주소값: %p\n", &data.value)
	data.value = "temp"
	data.score[3] = 100
}
