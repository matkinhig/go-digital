package db

func GetAllUsers() (*[]types.EUser, error) {
	db, err := config.InitDB()
	if err != nil {
		panic(err)
	}
	row, err := db.Querry
	sUsers := []types.EUser{}
}
