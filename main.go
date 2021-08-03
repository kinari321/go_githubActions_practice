package main

import (
	"fmt"
)

func hello() string {
	return "Hello, World!!!!"
}

func square(a, b int) int {
	return a * b
}

func main() {
	fmt.Println(hello())
	fmt.Println(square(1, 2))
}
