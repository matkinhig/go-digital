package config

import (
	"database/sql"
	_ "github.com/mattn/go-oci8"
)


func InitDB() (*sql.DB, error) {
	db, err := sql.Open("oci8", Config.OracleDB.Username+"/"+Config.OracleDB.Password+"@"+Config.OracleDB.Host)
	if err != nil {
		panic(err)
	}
	return db, err
}
