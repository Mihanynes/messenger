package models

import (
	"errors"
	"html"
	"strings"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"schedule/utils/token"
)

type User struct {
	gorm.Model
	Username  string `gorm:"size:255;not null;unique" json:"username"`
	Password  string `gorm:"size:255;not null;" json:"password"`
	Email     string `gorm:"size:255;not null;" json:"email"`
	FirstName string `gorm:"size:255;not null;" json:"first_name"`
	LastName  string `gorm:"size:255;not null;" json:"last_name"`
	ImageSrc  string `gorm:"size:255;not null;default:avatars/default.jpg" json:"image_src"`
}

func GetUserByID(uid uint) (User, error) {

	var u User

	if err := DB.First(&u, uid).Error; err != nil {
		return u, errors.New("User not found!")
	}

	u.PrepareGive()

	return u, nil

}

func (u *User) PrepareGive() {
	u.Password = ""
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(username string, password string) (string, error) {

	var err error

	u := User{}

	err = DB.Model(User{}).Where("username = ?", username).Take(&u).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := token.GenerateToken(u.ID)

	if err != nil {
		return "", err
	}

	return token, nil

}

func (u *User) SaveUser() (*User, error) {

	var err error
	err = DB.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) BeforeSave() error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	return nil

}

func FindUser(username string) ([]User, error) {

	var users []User
	if err := DB.Where("username LIKE ?", username+"%").Find(&users); err != nil {
		return users, err.Error
	}
	// Убрать пароли у каждого юзера
	return users, nil
}

func FindOneUser(username string) (User, error) {

	var user User
	if err := DB.Where(User{Username: username}).Find(&user); err != nil {
		return user, err.Error
	}
	// Убрать пароли у каждого юзера
	return user, nil
}

func (u *User) UpdatePhoto(imageSrc string) (*User, error) {
	if err := DB.Model(&u).Update("image_src", imageSrc).Error; err != nil {
		return u, err
	}
	return u, nil
}
