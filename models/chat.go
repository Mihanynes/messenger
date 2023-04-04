package models

import "github.com/jinzhu/gorm"

type Chat struct {
	gorm.Model
	FirstUserID  uint `gorm:"not null" json:"first_user_id"`
	SecondUserID uint `gorm:"not null" json:"second_user_id"`
}

func (chat *Chat) Create() (*Chat, error) {
	err := DB.Create(&chat).Error
	if err != nil {
		return chat, err
	}
	return chat, nil

}
