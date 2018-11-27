// Flagを使用し、コマンドオプションに渡す。
package main

import (
	"flag"
	"fmt"
	"os/exec"
)

func main() {
	var (
		cmd = flag.String("c", "default", "string flag")
		opt = flag.String("o", "default", "string flag")
	)
	flag.Parse()
	fmt.Println("cmd:", *cmd)
	fmt.Println("opt:", *opt)

//	out, err := exec.Command("*cmd", "*opt").Output()
	out := exec.Command(*cmd, *opt).Output()
//		if err != nil {
//		fmt.Println("Command exec error")
//	}
		fmt.Printf("string(out)")
}

