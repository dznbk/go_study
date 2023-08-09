package main

import "fmt"

// iotaは都度評価される
// Appleは 0 + 0
// Orangeは 1 + 5
// Bananaは 2 + 10
const (
	Apple = iota + iota
	Orange = iota + 5
	Banana = iota + iota + 10
)

func main() {
	fmt.Println(Apple)
	fmt.Println(Orange)
	fmt.Println(Banana)

	fmt.Println("Hello, World!")
}