package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	users := make(map[int]User)

	users[99] = User{"coco", 25}
	users[11] = User{"jin", 37}
	users[28] = User{"lily", 28}
	users[19] = User{"rora", 21}
	users[192] = User{"kate", 34}

	for key, value := range users {
		fmt.Println(key, value)
	}

}
