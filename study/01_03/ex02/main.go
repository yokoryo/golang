// [http://localhost:8000/hello]のときは「hello web server!」と表示
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello web server!\n")
	})

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}