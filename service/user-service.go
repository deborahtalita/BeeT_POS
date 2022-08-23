package service

import (
	"beet_pos/dto"
	"beet_pos/entity"
	"beet_pos/repository"
	"log"

	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	VerifyCredential(username string, password string, outlet_id string) interface{}
	IsDuplicate(username string) bool
	CreateUser(user dto.RegisterUserDTO) entity.User
	ReadUser() []entity.User
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}
func (service *userService) CreateUser(user dto.RegisterUserDTO) entity.User{
	userToCreate := entity.User{}
	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))
	if err != nil{
		log.Fatalf("Failed map %v",err)
	}

	res := service.userRepository.InsertUser(userToCreate)
	return res
}

// VerifyCredential implements UserService
func (svc *userService) VerifyCredential(username string, password string, outlet_id string) interface{} {
	log.Printf("UserService : VerifyCredential")
	res := svc.userRepository.VerifyCredential(username, password)
	if v, ok := res.(entity.User); ok {
		comparedPassword := comparePassword(v.User_password, []byte(password))
		if v.User_name == username && comparedPassword {
			// if v.Outlet_id == outlet_id {
				return res
			// }
		}
	}
	return false
}

func(service *userService) ReadUser()[]entity.User{
	return service.userRepository.ReadUser()
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

func (service *userService) IsDuplicate(username string) bool{
	res := service.userRepository.IsDuplicate(username)
	return !(res.Error == nil)
}



