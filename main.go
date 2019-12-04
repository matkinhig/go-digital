package main

import (
	"fmt"

	"github.com/matkinhig/go-digital/config"
	"github.com/matkinhig/go-digital/db"
	"github.com/go-xorm/xorm"
)

// type User struct {
//     Id int64
//     Name string
//     Salt string
//     Age int
//     Passwd string `xorm:"varchar(200)"`
//     Created time.Time `xorm:"created"`
//     Updated time.Time `xorm:"updated"`
// }

func main() {
	fmt.Println("Start Golang")
	fmt.Println(config.Config)
	db.MappingUser()
	// us := db.GetAllUsers()
	// fmt.Println(us)
	testXORM()
}

func testXORM() {
	engine , err = xorm.NewEngine("oci8",config.Config.OracleDB.Username+"/"+config.Config.OracleDB.Password+"@"+config.Config.OracleDB.Host)
	if err != nil {
		panic(err)
	}
	results, err := engine.Query("select * from ibkeusers")
	if err != nil {
		panic(err)
	}
	fmt.Println(results)

}

