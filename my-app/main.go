package main

import (
	"fmt"
	"github.com/mattn/go-runewidth"
)

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

	arrayStudy()

	fmt.Println("Hello, World!")
	fmt.Println(runewidth.StringWidth("こにゃにゃちわわあ"))
}

func arrayStudy() {
	a := make([]int, 5)
	fmt.Println(len(a))
	fmt.Println(cap(a))

	for i := 0; i < 5; i++ {
		a = append(a, i)
	}

	fmt.Println(len(a))
	fmt.Println(cap(a))
}