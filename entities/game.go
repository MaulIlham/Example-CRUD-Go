package entities

import (
	"github.com/go-ozzo/ozzo-validation"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

type Game struct {
	Id uuid.UUID `gorm:"type:varchar(255);primary_key;not null" json:"id"`
	Title string `gorm:"type:varchar(50);not null" json:"title"`
	Publisher string `gorm:"type:varchar(50);not null" json:"publisher"`
	Price string `gorm:"type:varchar(15);not null" json:"price"`
	ReleaseDate string `gorm:"type:varchar(12)" json:"release_date"`
	//foreignkey:'CategoryId' nama field

}

func (g Game) BeforeCreate(scope *gorm.Scope) error  {
	return scope.SetColumn("Id",uuid.NewV4())
}

func (g Game) Validate() error  {
	return validation.ValidateStruct(&g,
		validation.Field(&g.Title, validation.Required.Error("This field still empty")),
		validation.Field(&g.Price, validation.Required.Error("This field still empty")),
		validation.Field(&g.Publisher, validation.Required.Error("This field still empty")),
		validation.Field(&g.ReleaseDate, validation.Required.Error("This field still empty")),
	)
}