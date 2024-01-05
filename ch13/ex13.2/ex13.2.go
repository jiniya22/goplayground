package main

import (
	"fmt"
	"unsafe"
)

type Student struct {
	Age   int32
	F1    int8
	F2    int8
	Score float64
}

func main() {

	u1 := Student{Age: 32, Score: 9.5}
	fmt.Println(unsafe.Sizeof(u1), &u1.Age, &u1.F1, &u1.F2, &u1.Score)

	fmt.Printf("%d, %d, %d, %d, %d\n", unsafe.Sizeof(u1), &u1.Age, &u1.F1, &u1.F2, &u1.Score)

	//u1 := Student{Age: 32, Score: 9.5}
	//fmt.Println(unsafe.Sizeof(u1), &u1.Age, &u1.Score)
	//fmt.Printf("%d, %d, %d\n", unsafe.Sizeof(u1), &u1.Age, &u1.Score)

}
