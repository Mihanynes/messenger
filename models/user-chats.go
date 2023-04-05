package models

func (u *User) GetAllUserChats() ([]Chat, error) {
	var chats []Chat
	err := DB.Where(&Chat{FirstUserID: u.ID}).Or(&Chat{SecondUserID: u.ID}).Find(&chats)
	if err != nil {
		return chats, err.Error
	}
	return chats, nil
}

type ChatIcon struct {
	user    User
	chat    Chat
	message Message
}

//работает, но не корректо вовзращает json

func (u *User) GetLastMessages() ([]ChatIcon, error) {
	var chatIcons []ChatIcon
	chatIcons = nil
	chats, err := u.GetAllUserChats()
	for _, chat := range chats {
		var _user User
		if chat.FirstUserID != u.ID {
			_user, err = GetUserByID(chat.FirstUserID)

		} else {
			_user, err = GetUserByID(chat.SecondUserID)
		}
		_chat := chat
		_message, _ := chat.GetLastMessage()
		chatIcon := ChatIcon{_user, _chat, _message}
		chatIcons = append(chatIcons, chatIcon)
	}
	if err != nil {
		return chatIcons, err
	}
	return chatIcons, nil
}
