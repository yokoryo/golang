package main

//gogueryを使ってみた
import (
	"fmt"
	"net/http"
	"github.com/PuerkitoBio/goquery"
)

func get(url string) {
	doc, _ := goquery.NewDocument(url)
	doc.Find("a").Each(func(_ int, s *goquery.Selection) {
			 url, _ := s.Attr("href")
			 fmt.Println(url)
	})
}

func main() {
	fs := http.FileServer(http.Dir("public/"))
	http.Handle("/", http.StripPrefix("/", fs))
	http.ListenAndServe(":9999", nil)
	url := "http://localhost:9999/"
	get(url)
}