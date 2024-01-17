package main

import (
	"fmt"
	"regexp"
)

func main() {
	//f1("483")
	//f1("42.83")
	//f1("rk483")
	//f1("12.3456789")
	//f1("123456789")
	fmt.Println(682 % 500)
}

func f1(s string) {
	if IsInt(s) {
		fmt.Printf("%v, 정수, ", s)
	} else {
		fmt.Printf("%v, 정수 x, ", s)
	}
	if IsFloat(s) {
		fmt.Println("실수")
	} else {
		fmt.Println("실수 x")
	}
}

func IsInt(s string) bool {
	return regexp.MustCompile("^\\d+$").MatchString(s)
}

func IsFloat(s string) bool {
	return regexp.MustCompile("^\\d*[.]\\d*$").MatchString(s)
}
