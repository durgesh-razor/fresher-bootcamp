package main

import (
	"studentRecord/Config"
	"studentRecord/Modules/Marks"
	"studentRecord/Modules/Student"
	"studentRecord/Routes"
)

func main() {
	DB, err := Config.ConnectDB()
	if err != nil {
		panic(err)
		// fmt.Println("Status:", err)
	}

	//Apply Migrations
	DB.Migrator().DropTable(&Student.Student{}, &Marks.Marks{})
	DB.AutoMigrate(&Student.Student{}, &Marks.Marks{})
	r := Routes.SetupRouter()
	//running
	r.Run()
}
