package main

import (
	"docker_sample/database"
	"fmt"
)

type User struct {
	Id    int    `json:"id" param:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	db := database.Connect()
	defer db.Close()

	err := db.Ping()

	if err != nil {
		fmt.Println(err.Error())
		return
	} else {
		fmt.Println("データベース接続成功")
	}
}
