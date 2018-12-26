//WEBページのhttpレスポンスコードを表示
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	)

func main() {
	//url := flag.String("url", "http://localhost:8000", "URL")
	url := flag.String("url", "https://golang.org", "URL")
	flag.Parse()

	res, err := http.Get(*url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	
	status := res.StatusCode
	//if err != nil {
	//	log.Fatal(err)
	//}
	fmt.Printf("%d\n", status)
}