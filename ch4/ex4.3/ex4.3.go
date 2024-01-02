package main

import "fmt"

func main() {
	a := 5
	var b float32 = 3.14

	var c = int8(a) * 3
	d := int(b)
	e := 10 * int64(b)
	f := 5 * 7

	fmt.Printf("a = %v (%T)\n", a, a)
	fmt.Printf("b = %v (%T)\n", b, b)
	fmt.Printf("c = %v (%T)\n", c, c)
	fmt.Printf("d = %v (%T)\n", d, d)
	fmt.Printf("e = %v (%T)\n", e, e)
	fmt.Printf("f = %v (%T)\n", f, f)
}
