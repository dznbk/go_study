package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sync"
	"csv"
	"bytes"
)

func downloadCSV(wg *sync.WaitGroup, urls []string, ch chan []byte) {
	defer wg.Done()
	defer close(ch)

	// ダウンロード処理
	for _, u := range urls {
		resp, err := http.Get(u)
		if err != nil {
			log.Println("cannot download CSV:", err)
			continue
		}
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			resp.Body.Close()
			log.Println("cannot read content:", err)
			continue
		}
		resp.Body.Close()
		ch <- b
	}
}

func main() {
	urls := []string{
		"https://example.com/1.csv",
		"https://example.com/2.csv",
		"https://example.com/3.csv",
		"https://example.com/4.csv",
		"https://example.com/5.csv",
	}
	// バイト列を転送するためのチャンネル1
	ch := make(chan []byte)

	var wg sync.WaitGroup
	wg.Add(1)
	go downloadCSV(&wg, urls, ch)

	// goroutineからコンテンツ受け取る
	for _, b := range ch {
		r := csv.NewReader(bytes.NewReader(b))
		for {
			records, err := r.Read()
			if err != nil {
				log.Fatal(err)
			}
			insertRecords(records)
		}
	}
	wg.Wait()
}