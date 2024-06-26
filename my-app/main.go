package main

import (
	"fmt"
)

// iotaは都度評価される
// Appleは 0 + 0
// Orangeは 1 + 5
// Bananaは 2 + 10
const (
	Apple  = iota + iota
	Orange = iota + 5
	Banana = iota + iota + 10
)

var name = "dpon"

func init() {
	fmt.Println("init " + name)
}

func main() {
	fmt.Println(Apple)
	fmt.Println(Orange)
	fmt.Println(Banana)

	arrayStudy()

	fmt.Println("Hello! " + name)
	fmt.Println("Hello, World!")
	fmt.Println(RX78)
	fmt.Println(zeta)
}

func arrayStudy() {
	a := make([]int, 5)
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))

	for i := 0; i < 5; i++ {
		a = append(a, i)
	}

	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
}
