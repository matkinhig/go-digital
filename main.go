package main

import (
	"fmt"

	"github.com/matkinhig/go-digital/config"
	"github.com/matkinhig/go-digital/db"
)

func main() {
	fmt.Println("Start Golang")
	fmt.Println(config.Config)
	db.MappingUser()
	// us := db.GetAllUsers()
	// fmt.Println(us)
}
