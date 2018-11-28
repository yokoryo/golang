// 1行のtxt形式のファイルを読み取り、内容すべて表示
package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {

	cmd := exec.Command("docker", *opt)
	result, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
		fmt.Printf("出力結果: \n%s", result)
}

