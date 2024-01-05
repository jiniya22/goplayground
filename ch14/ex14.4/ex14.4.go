package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	var u User
	var p1 = &u
	fmt.Println(*p1)

	var p2 = &User{}
	fmt.Println(*p2)

	var p3 = new(User)
	fmt.Println(*p3)

	var a1 = &User{Name: "jini"}
	var a2 = a1
	var a3 = a2
	fmt.Println(a1, &a1.Name)
	fmt.Println(a2, &a2.Name)
	fmt.Println(a3, &a3.Name)
}
