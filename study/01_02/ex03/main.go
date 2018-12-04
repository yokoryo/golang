// 1行のjson形式のファイルを読み取り、内容すべてを表示
package main

import (
	"fmt"
	"os"
	"encoding/json"
	"io/ioutil"
)

typ Person struct {
	id int 'json:"id"'
	name string 'jdon:"name"'
	age int 'json:"age"'
	ability struct {
		programing string 'json:"programing"'
		operation string 'json:"operation"'
	} 'json:"vivd_info"'
}

func main() {
	func Unmarshal(data []byte, v interface{}) error
	var persons []Person
	for _, p := range persons {
		fmt.Printf("%d : %s : %s : %s : %s\n", p.ID, p.Name, p.Birthday, p.VividInfo.Color, p.VividInfo.Weapon)
	}
}