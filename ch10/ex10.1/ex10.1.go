package main

import "fmt"

func main() {
	day := "friday"

	switch day {
	case "friday":
		fmt.Println("금요일은 불금!")
	case "saturday", "sunday":
		fmt.Println("주말 너무 좋아요")
	default:
		fmt.Println("월화수목 월화수목")
	}
}
