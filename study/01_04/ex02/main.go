package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	)

func query() {
	res, err := http.Get("http://localhost:9999")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%q\n", body[:])
}

func main() {
	fs := http.FileServer(http.Dir("public/"))
	http.Handle("/", http.StripPrefix("/", fs))
	go http.ListenAndServe(":9999", nil)
	go query()
	time.Sleep(time.Second)
}