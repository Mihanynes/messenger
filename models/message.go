package models

import "github.com/jinzhu/gorm"

type Message struct {
	gorm.Model
	SenderID uint   `gorm:"not null" json:"sender_id"`
	ChatID   uint   `gorm:"not null" json:"chat_id"`
	Text     string `gorm:"size:255;not null;" json:"text"`
}

func (message *Message) Create() (*Message, error) {
	err := DB.Create(&message).Error
	if err != nil {
		return message, err
	}

	var chat Chat
	DB.First(&chat, message.ChatID)
	if err := DB.Model(&chat).Update("last_message", message.Text).Error; err != nil {
		return message, err
	}
	return message, nil
}

func (message *Message) Update() (*Message, error) {
	if err := DB.Model(&message).Update("text", message.Text).Error; err != nil {
		return message, err
	}
	return message, nil
}
