package service

import (
	"beet_pos/entity"
	"beet_pos/repository"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	VerifyCredential(username string, password string, outlet_id string) interface{}
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

// VerifyCredential implements UserService
func (svc *userService) VerifyCredential(username string, password string, outlet_id string) interface{} {
	log.Printf("UserService : VerifyCredential")
	res := svc.userRepository.VerifyCredential(username, password)
	v := res.(entity.User)

	if (v.Username == username) && (v.Password == password){
		if (v.Outlet_ID == outlet_id){
			return res
		}
	}
	return false
}

func comparePassword(hashedPass string, plainPass []byte) bool {
	byteHash := []byte(hashedPass)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPass)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}