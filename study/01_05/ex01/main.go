package main

import (
	"database/sql"
	//"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//接続
	//db, err := sql.Open("mysql", "go_user:@go_apps")
	goconf := "go_user:@/tcp(:3306)go_apps"
	db, err := sql.Open("mysql", goconf)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	
	/*
	//Select
	rows, err := db.Query("select * from users")
	defer rows.Close()
	if err != nil {
			panic(err.Error())
	}

	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
				panic(err.Error())
		}
		fmt.Println(id, name)
	}
	*/
}