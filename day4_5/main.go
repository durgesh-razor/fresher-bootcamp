package main

import (
	"retailer-service/Config"
	"retailer-service/Models"
	"retailer-service/Router"
)

var err error

func main() {
	Config.DB, err = Config.ConnectDB()
	if err != nil {
		panic(err)
	}

	//Apply Migrations
	// Config.DB.Migrator().DropTable(&Models.Orders{}, &Models.Product{}, &Models.User{})
	Config.DB.AutoMigrate(&Models.Orders{}, &Models.Product{}, &Models.User{})
	r := Router.SetupRouter()
	//running
	r.Run()
}
