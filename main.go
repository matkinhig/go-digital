package main

import (
	"fmt"
	"github.com/matkinhig/go-digital/config"
	"github.com/matkinhig/go-digital/db"
)

func main() {
	fmt.Println("Start Golang")
	fmt.Println(config.Config)
	s := db.GetAllUsers()
	fmt.Println(s)
}
