package main

import (
	"fmt"
	"strings"
)

func main() {
	a := 1
	for a < 5 {
		fmt.Println("Number:", a)
		a++
	}
	fmt.Println(strings.ToUpper("Hello"))
}
