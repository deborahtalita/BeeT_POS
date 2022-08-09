package repository

import (
	"beet_pos/structs"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)


type UserRepository interface {
	InsertUser(user structs.User) structs.User
}

type userConnection struct{
	connection *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository{
	return &userConnection{
		connection: db,
	}
}

func (db *userConnection) InsertUser(user structs.User) structs.User{
	user.User_password = hashAndSalt([]byte(user.User_password))
	db.connection.Save(&user)
	return user

}

func hashAndSalt(pwd []byte) string{
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil{
		log.Println(err)
		panic("Failed to hash password")
	}
	return string(hash)
}
