package models

import "github.com/jinzhu/gorm"

func (u *User) GetAllUserChats() ([]Chat, error) {
	var chats []Chat
	err := DB.Where(&Chat{FirstUserID: u.ID}).Or(&Chat{SecondUserID: u.ID}).Find(&chats)
	if err != nil {
		return chats, err.Error
	}
	return chats, nil
}

type ChatIcon struct {
	gorm.Model
	User    User    `gorm:"not null" json:"user"`
	Chat    Chat    `gorm:"not null" json:"chat"`
	Message Message `gorm:"not null" json:"message"`
}

//работает, но не корректо вовзращает json

func (u *User) GetLastMessages() ([]ChatIcon, error) {
	var chatIcons []ChatIcon
	chats, err := u.GetAllUserChats()
	for _, chat := range chats {
		var user User
		if chat.FirstUserID != u.ID {
			user, err = GetUserByID(chat.FirstUserID)

		} else {
			user, err = GetUserByID(chat.SecondUserID)
		}
		var chatIcon ChatIcon
		chatIcon.User = user
		chatIcon.Chat = chat
		chatIcon.Message, _ = chat.GetLastMessage()
		chatIcons = append(chatIcons, chatIcon)
	}
	if err != nil {
		return chatIcons, err
	}
	return chatIcons, nil
}
