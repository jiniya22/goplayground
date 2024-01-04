package main

import "fmt"

type Member struct {
	Name string
	Age  int
}

func main() {
	m := Member{Name: "jay", Age: 38}
	fmt.Println(m)

	m1 := Member{"jini", 32}
	fmt.Println(m1)

	var m2 Member
	m2.Name = "coco"
	m2.Age = 22
	fmt.Println(m2)

	m3 := Member{
		"rora",
		29,
	}
	fmt.Println(m3)
}
