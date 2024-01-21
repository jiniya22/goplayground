package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["apple"] = 3
	m["banana"] = 12
	m["melon"] = 22

	delete(m, "banana")
	delete(m, "tomato")
	fmt.Println("있는 요소를 조회할 경우: ", m["apple"])
	fmt.Println("없는 요소를 조회할 경우: ", m["banana"])

	if value, ok := m["grape"]; ok {
		fmt.Println("저장된 값은", value)
	} else {
		fmt.Println("해당 값이 존재하지 않습니다.")
	}
}
