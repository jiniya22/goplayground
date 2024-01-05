package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	pUser := NewUser("jini", 32)
	fmt.Println(*pUser)
}

func NewUser(name string, age int) *User {
	return &User{Name: name, Age: age}
}
