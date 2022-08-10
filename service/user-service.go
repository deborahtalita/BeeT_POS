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
	if v, ok := res.(entity.User); ok {
		comparedPassword := comparePassword(v.User_password, []byte(password))
		if v.User_name == username && comparedPassword {
			if v.Outlet_id == outlet_id {
				return res
			}
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