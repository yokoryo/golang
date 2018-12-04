// 1行のtxt形式のファイルを読み取り、内容すべて表示
package main

import (
	"fmt"
	"os"
)

const BUFSIZE = 1024

func main() {
	file, err := os.Open("./sample.txt")
	if err != nil {
		fmt.Println("Open error")
	}
	defer file.Close()

	put := make([]byte, BUFSIZE)
	for {
		num, err := file.Read(put)
		if num == 0 {
			break
		}
		if err != nil {
			fmt.Println("Read error")
			break
		}

		fmt.Println(string(put[:num]))
	}
}