package main

import (
	"gamehsop/controllers"
	"gamehsop/entities"
	"gamehsop/repositories"
	"log"

)

func main()  {
	db, err := repositories.Connect()
	if err != nil {
		log.Println("Connection Failed!")
		panic("failed")
	}
	defer db.Close()
	db.AutoMigrate(&entities.Game{})
	controllers.HandleRequest()
}
