package db

import (
	"fmt"
	"log"

	"github.com/matkinhig/go-digital/config"
	"github.com/matkinhig/go-digital/types"
)

func GetAllUsers() (*[]types.EUsers, error) {
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
	return &sUsers, nil
}
