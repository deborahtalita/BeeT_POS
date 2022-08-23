package repository

import (
	"beet_pos/entity"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
	VerifyCredential(username string, password string) interface{}
	InsertUser(user entity.User) entity.User
	IsDuplicate(username string) (tx *gorm.DB)
	ReadUser() []entity.User
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
	log.Printf("UserRepository: VerifyCredential")
	var user entity.User
	res := db.connection.Where("user_name = ?", username).Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}

func (db *userConnection) InsertUser(user entity.User) entity.User{
	user.User_password = hashAndSalt([]byte(user.User_password))
	db.connection.Save(&user)
	return user

}

func (db *userConnection) ReadUser() []entity.User{
	var user []entity.User
	db.connection.Preload("Outlet").Find(&user)
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

func (db *userConnection) IsDuplicate(username string) (tx *gorm.DB){
	var user entity.User
	return db.connection.Where("user_name = ?", username).Take(&user)
}
