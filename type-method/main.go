package main

import "fmt"

type Value int

// typeはメソッドを生やせる
// vはレシーバ
func (v *Value) Add(n Value) {
	*v += n
}

func main() {
	v := Value(1)
	v.Add(3)
	fmt.Println(v)
}