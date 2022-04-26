package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Name string `json:"name"`
	UserName string `json:"user_name" gorm:"unique"`
	Email string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

func (u *User) Hashpassword(password string)(error)  {
	generateFromPassword, err := bcrypt.GenerateFromPassword([]byte(password),14)
	if err != nil {
		log.Fatalln("Error in bcrypting password")
		return err
	}
	u.Password = string(generateFromPassword)
	return nil
}

func (u User) Checkpassword(encryptedPassword string)(error)  {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(encryptedPassword))
	if err != nil {
		return err
	}
	return nil
}
