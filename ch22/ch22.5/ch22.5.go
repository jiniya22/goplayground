package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["jini"] = 100
	m["coco"] = 92
	m["sol"] = 84

	for key, value := range m {
		fmt.Printf("key: %s, value: %d\n", key, value)
	}
	fmt.Println(m["jini"])
}
