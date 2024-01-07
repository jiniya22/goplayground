package main

import (
	"ch16/ex16.2/publicpkg"
	"fmt"
)

func main() {
	fmt.Printf("PI: %v, ScreenSize: %v\n", publicpkg.PI, publicpkg.ScreenSize)
	publicpkg.PublicFunc()

	var myInt publicpkg.MyInt = 10
	fmt.Println("myInt:", myInt)

	myStruct := publicpkg.MyStruct{Name: "jini"}
	fmt.Println(myStruct)
	myStruct.PublicMethod()
}
