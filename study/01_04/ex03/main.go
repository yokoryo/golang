package main

import (
	"fmt"
	"net/http"
	"github.com/PuerkitoBio/goquery"
)

var cnt int = 0
func query(url string) {
	doc, _ := goquery.NewDocument(url)
	doc.Find("a").Each(func(_ int, s *goquery.Selection) {
			 cnt += 1
	})
}	

func main() {
	fs := http.FileServer(http.Dir("public/"))
	http.Handle("/", http.StripPrefix("/", fs))
	go http.ListenAndServe(":9999", nil)
	url := "http://localhost:9999/"
	query(url)
	fmt.Printf("num of tag : %d\n", cnt)
}