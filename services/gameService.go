package services

import (
	"gamehsop/entities"
	"gamehsop/repositories"
)

func GetAllDataGame() (*[]entities.Game, error) {
	db, err := repositories.Connect()
	if err != nil {
		return nil, err
	}
	games := []entities.Game{}
	db.Find(&games)
	return &games, nil
}

func SearchDataGame(title string) (*[]entities.Game, error)  {
	db, err := repositories.Connect()
	if err != nil {
		return nil, err
	}
	games := []entities.Game{}
	db.Where("title like ?","%"+title+"%").Find(&games)
	return &games, nil
}

func GetDataGameById(id string) (*entities.Game, error)  {
	db, err := repositories.Connect()
	if err != nil {
		return nil, err
	}
	var result entities.Game
	db.Where("id=?",id).Find(&result)
	return &result, nil
}

func SaveDataGame(game entities.Game) (*entities.Game,error) {
	db, err := repositories.Connect()
	if err != nil {
		return nil,err
	}
	error := game.Validate()
	if error != nil {
		return nil, error
	}
	db.Create(&game)
	return &game,nil
}

func UpdateDataGame(game entities.Game) error {
	db, err := repositories.Connect()
	if err != nil {
		return err
	}
	var newGame entities.Game
	db.Where("id=?",game.Id).Find(&newGame)
	newGame = game
	return db.Save(&newGame).Error
}

func DeleteDataGame(id string) error  {
	db, err := repositories.Connect()
	if err != nil {
		return err
	}
	var game entities.Game
	db.Where("id=?",id).Find(&game)
	return db.Delete(&game).Error
}