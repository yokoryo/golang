// 1行のjson形式のファイルを読み取り、内容すべてを表示
package main

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"os"
)

type Person struct {
	Id		int		`json:"id"`
	Name	string	`json:"name"`
	Age		string	`json:"age"`
	Ability struct {
		Pro	string	`json:"programming"`
		Ope	string	`json:"operation"`
	}
}

func main() {
	file, err := ioutil.ReadFile("./sample.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "file open err!: %v\n", err)
		os.Exit(1)
	}
	
	var persons []Person

	if err := json.Unmarshal(file, &persons); err != nil {
		fmt.Fprintf(os.Stderr, "json unmarshal err!: %v\n", err)
		os.Exit(1)
	}
	
	//fmt.Println(string(file))
	for _, p := range persons {
		fmt.Printf("%d : %s : %s : %s : %s\n", p.Id, p.Name, p.Age, p.Ability.Pro, p.Ability.Ope)
	}
}