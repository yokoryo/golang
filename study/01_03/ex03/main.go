package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
)

func main() {
	res, err := http.Get("https://golang.org")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b))
}