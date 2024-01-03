package main

import "fmt"

type ColorType int

const (
	Red ColorType = iota
	Blue
	Yellow
	Black
)

func colorToString(color ColorType) string {
	switch color {
	case Red:
		return "red"
	case Blue:
		return "blue"
	default:
		return "etc"
	}
}

func getMyColor() ColorType {
	return Blue
}

func main() {
	fmt.Println("내가 가장 좋아하는 색은", colorToString(getMyColor()))
}
