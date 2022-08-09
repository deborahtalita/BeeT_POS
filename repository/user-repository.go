package repository

import (
	"beet_pos/entity"
	"log"

	"gorm.io/gorm"
)

type UserRepository interface {
	VerifyCredential(username string, password string) interface{}
}

type userConnection struct {
	connection *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}


func (db *userConnection) VerifyCredential(username string, password string) interface{} {
	var user entity.User
	res := db.connection.Where("username = ?", username).Take(&user)
	log.Printf(username)
	if res.Error == nil {
		return user
	}
	return nil
}
