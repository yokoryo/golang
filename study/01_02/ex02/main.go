// 複数行のtxt形式のファイルを読み取り、1行ずつ1秒間隔で表示
package main

import (
	"fmt"
	"os"
	"bufio"
)

const BUFSIZE = 1024

func main() {
	file, err := os.Open("./sample.txt")
	if err != nil {
		fmt.Println("Open error")
	}
	defer file.Close()

	//1行ずつ読み取り、秒間隔ではない。。。
	put := bufio.NewScanner(file)
	for put.Scan(){
		line := put.Text()
		fmt.Println(line)
	}
}