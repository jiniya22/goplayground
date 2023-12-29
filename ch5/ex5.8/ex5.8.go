package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)
	var a, b int

	n, err := fmt.Scan(&a, &b)
	if err != nil {
		fmt.Println(n, err)
		stdin.ReadString('\n')
	} else {
		fmt.Println(n, a, b)
	}

	n, err = fmt.Scan(&a, &b)
	if err != nil {
		fmt.Println(n, err)
		stdin.ReadString('\n')
	} else {
		fmt.Println(n, a, b)
	}
}
