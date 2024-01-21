package main

import (
	"errors"
	"fmt"
	"math"
)

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		//return 0, fmt.Errorf("제곱근은 양수여야 합니다. 입력값: %f\n", f)
		return 0, errors.New("제곱근은 양수여야 합니다")
	}
	return math.Sqrt(f), nil
}

func f1(num float64) {
	result, err := Sqrt(num)
	if err != nil {
		fmt.Printf("Error! %v\n", num)
		return
	}
	fmt.Printf("Sqrt(%f) = %v\n", num, result)
}

func main() {
	f1(2.2)
	f1(-2.2)
}
