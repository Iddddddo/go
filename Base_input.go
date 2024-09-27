package main

import "fmt"

var (
	name string
)

func main() {
	fmt.Println("Как тебя зоут?")
	fmt.Scanf("%s\n", &name)
	fmt.Printf("Приятно познакомиться, %s", name)
}
