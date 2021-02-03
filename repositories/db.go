package repositories

import (
	"errors"
	"fmt"
	"gamehsop/utils"
	"github.com/jinzhu/gorm"
)

var (
	DB_USER,
	DB_PASSWORD,
	DB_HOST,
	DB_PORT,
	SCHEMA_NAME string
)

func Connect() (*gorm.DB, error) {
	DB_USER = utils.GetEnv("DB_USER","admin")
	DB_PASSWORD = utils.GetEnv("DB_PASSWORD","P@ssword123")
	DB_HOST = utils.GetEnv("DB_HOST","localhost")
	DB_PORT = utils.GetEnv("DB_PORT","3306")
	SCHEMA_NAME = utils.GetEnv("SCHEMA_NAME","gameshop?charset=utf8&parseTime=True")
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",DB_USER,DB_PASSWORD,DB_HOST,DB_PORT,SCHEMA_NAME)
	db, err := gorm.Open("mysql", dataSourceName)
	if err != nil {
		return nil, errors.New("Connection Failed to Open")
	}
	return db, nil
}
