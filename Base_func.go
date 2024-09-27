package main

import "fmt"

var (
	a, b int
)

func def(a, b int) (c int) {
	c = a + b
	return c
}

func main() {
	fmt.Scan(&a, &b)
	fmt.Println(def(a, b))
}
