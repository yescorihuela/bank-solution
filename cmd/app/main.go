package main

import "github.com/yescorihuela/bluesoft-bank-solution/internal/shared/utils"

func main() {
	_, err := utils.LoadConfig("../../")
	if err != nil {
		panic(err)
	}

	/*
		db, err := databases.NewPostgresqlDbConnection(config)
		if err != nil {
			panic(err)
		}
		defer db.Close()
	*/

}
