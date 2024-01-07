package main

import (
	"fmt"
	"github.com/guptarohit/asciigraph"
	"goproject/usepkg/custompkg"
)

func main() {
	custompkg.PrintLine()
	data := []float64{3, 4, 5, 8, 10, 5, 2}
	graph := asciigraph.Plot(data)
	fmt.Println(graph)
}
