package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// URLを引数で指定し、実行する場合
func cli() {
	urlarg := flag.String("url", "https://github.com/n-guitar/go_study", "引数のURL書かないとか、まじか・・・")
	flag.Parse()
	req, _ := http.NewRequest("GET", *urlarg, nil)
	client := new(http.Client)
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	html, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(html))
}

func main() {
	fs := http.FileServer(http.Dir("public/"))
	http.Handle("/", http.StripPrefix("/", fs))

	go http.ListenAndServe(":9999", nil)
	go cli()
	time.Sleep(3 * time.Second)
}