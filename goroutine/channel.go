package main

import "fmt"

func server(ch chan string){
	defer close(ch)
	// ch <- はチャンネルへのデータ送信
	ch <- "one"
	ch <- "two"
	ch <- "three"
}

func main() {
	var s string

	// チャンネルの作成
	ch := make(chan string)
	go server(ch)
	// <- ch　はチャンネルからのデータ受信
	s = <- ch
	fmt.Println(s)
	s = <- ch
	fmt.Println(s)
	s = <- ch
	fmt.Println(s)
}
