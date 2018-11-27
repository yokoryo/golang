// Flagを使用し、コマンドオプションに渡す。
package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	opt := flag.String("o", "ps", "option")
	flag.Parse()

//	fmt.Println("opt: %s", *opt)

	cmd := exec.Command("docker", *opt)
	result, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
		fmt.Printf("出力結果: \n%s", result)
}

