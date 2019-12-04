package db

import (
	"context"
	"fmt"
	"log"

	"github.com/matkinhig/go-digital/config"
	"github.com/matkinhig/go-digital/types"
)

func GetAllUsers() *[]types.EUsers {
	fmt.Println("Start querry all users")
	db, err := config.InitDB()
	if err != nil {
		panic(err)
	}
	rows, err := db.Query("select * from halong.ibkeusers")
	if err != nil {
		panic(err)
	}
	var sUsers []types.EUsers
	for rows.Next() {
		s := types.EUsers{}
		err = rows.Scan(&s.EUSERID, &s.KHID, &s.TENTRUYCAP, &s.MATKHAU, &s.KHOA, &s.THOIGIANKHOA, &s.TRUYCAPGANNHAT, &s.DATELASTMAINT, &s.NOTE, &s.KHID_FBE)
		if err != nil {
			log.Fatal(err)
			panic(err)
		}
		sUsers = append(sUsers, s)
	}
	fmt.Println(sUsers[1])
	return &sUsers
}

func MappingUser() {
	fmt.Println("Start querry all users")
	db, err := config.InitDB()
	if err != nil {
		panic(err)
	}
	rows, err := db.QueryContext(context.TODO(), "select * from halong.ibkeusers")
	if err != nil {
		panic(err)
	}
	filter := map[string]interface{}{}
	for rows.Next() {
		err = rows.Scan(&filter)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println(filter)
}
