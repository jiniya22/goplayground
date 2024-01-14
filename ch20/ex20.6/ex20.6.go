package main

import "fmt"

type Shape interface {
	Introduce() string
}

type Polygon interface {
	Shape
	GetLines() int
}

type Circle struct {
	Name   string
	Radius int
}

type Square struct {
	Name  string
	Width int
}

func (v *Circle) Introduce() string {
	return fmt.Sprintf("%s는 원. 반지름은 %d 입니다.", v.Name, v.Radius)
}

func (v *Square) Introduce() string {
	return fmt.Sprintf("%s는 정사각형. 변의 길이는 %d 입니다.", v.Name, v.Width)
}

func (v *Square) GetLines() int {
	return 4
}

func PrintIntroduce(s Shape) {
	if p, ok := s.(Polygon); ok {
		fmt.Printf("변의 수는 %d\n", p.GetLines())
	}
	fmt.Println(s.Introduce())
}

func main() {
	var s1 Shape = &Circle{"원1", 2}
	fmt.Println(s1.Introduce())
	fmt.Printf("s1의 타입은 %T\n", s1)

	var s2 Shape = &Square{"사각형1", 5}
	fmt.Println(s2.Introduce())
	fmt.Printf("s2의 타입은 %T\n", s2)

	fmt.Println("-------")
	PrintIntroduce(s1)
	PrintIntroduce(s2)
}
