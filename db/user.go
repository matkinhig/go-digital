package db

import (
	"fmt"
	"log"

	"github.com/matkinhig/go-digital/config"
	"github.com/matkinhig/go-digital/types"
)

func GetAllUsers() (*[]types.EUsers{}, error) {
	fmt.Println("Start querry all users")
	db, err := config.InitDB()
	if err != nil {
		panic(err)
	}
	rows, err := db.Query("select * from halong.ibkeusers")
	if err != nil {
		panic(err)
	}
	sUsers := []types.EUsers{}
	for rows.Next() {
		var us types.EUsers
		err = rows.Scan(&us)
		if err != nil {
			log.Fatal(err)
			panic(err)
		}
		sUsers = append(sUsers, us)
	}
	fmt.Println(sUsers[1])
	return sUsers, err
}
