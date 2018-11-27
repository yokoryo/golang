// Linuxコマンドを実行する。
package main

import (
	"fmt"
	"os/exec"
)

func main() {
	out, err := exec.Command("ls", "-la").Output()	//cmd.Output()は標準出力を表示させる
	if err != nil {
		fmt.Println("Command exec error")
	}
		fmt.Printf(string(out))
}
