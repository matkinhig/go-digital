package config

import "database/sql"

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("oci8", Config.OracleDB.Username+"/"+Config.OracleDB.Password+"@"+Config.OracleDB.Host)
	if err != nil {
		panic(err)
	}
	return db, err
}
