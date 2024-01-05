package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(ToUpper("Hello World!"))
}

func ToUpper(s string) string {
	var builder strings.Builder
	for _, ch := range s {
		if ch >= 'a' && ch <= 'z' {
			builder.WriteRune('A' + (ch - 'a'))
		} else {
			builder.WriteRune(ch)
		}

	}
	return builder.String()
}
