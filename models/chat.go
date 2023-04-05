package models

import (
	"github.com/jinzhu/gorm"
)

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

func (chat *Chat) FindChat(first_user uint, second_user uint) (*Chat, error) {
	if first_user > second_user {
		tmp := first_user
		first_user = second_user
		second_user = tmp
	}
	err := DB.Where(&Chat{FirstUserID: first_user, SecondUserID: second_user}).First(&chat)
	if err != nil {
		chat.Create()
		return chat, nil
	}
	return chat, nil
}

func (chat *Chat) GetAllMessages() ([]Message, error) {
	var messages []Message
	err := DB.Where(&Message{ChatID: chat.ID}).Find(&messages)
	if err != nil {
		return messages, err.Error
	}
	return messages, nil
}

func (chat *Chat) GetLastMessage() (Message, error) {
	var message Message
	err := DB.Where(&Message{ChatID: chat.ID}).Last(&message)
	if err != nil {
		return message, err.Error
	}
	return message, nil
}
