package main

import (
	"studentRecord/Config"
	"studentRecord/Modules/Marks"
	"studentRecord/Modules/Student"
	"studentRecord/Routes"
)

var err error

func main() {
	Config.DB, err = Config.ConnectDB()
	if err != nil {
		panic(err)
	}

	//Apply Migrations
	Config.DB.Migrator().DropTable(&Student.Student{}, &Marks.Marks{})
	Config.DB.AutoMigrate(&Student.Student{}, &Marks.Marks{})
	r := Routes.SetupRouter()
	//running
	r.Run()
}
