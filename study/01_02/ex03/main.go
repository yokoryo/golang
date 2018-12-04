// 1行のjson形式のファイルを読み取り、内容すべてを表示
package main

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
)

type Person struct {
	Id		int			`json:"id"`
	Name	string	`json:"name"`
	Age		int			`json:"age"`
	Ability struct {
		pro	string	`json:"pro"`
		ope	string	`json:"ope"`
	}
}

func main() {
	file, err := ioutil.ReadFile("./sample.json")
	if err != nil {
		fmt.Println("Open error")
	}
	
	var persons []Person

	if err := json.Unmarshal(file, persons); err != nil {
		fmt.Printf("json unmarshal error")
	}
	

	for _, p := range persons {
		fmt.Printf("%d : %s : %d : %s : %s\n", p.Id, p.Name, p.Age, p.Ability.pro, p.Ability.ope)
	}
}