package service

import (
	"beet_pos/dto"
	"beet_pos/repository"
	"beet_pos/structs"
	"log"

	"github.com/mashingan/smapping"
)

type CustomerService interface {
	RegisterCustomer(customer dto.RegisterCustomerDTO) structs.Customer
	IsDuplicate(customer_phone string)bool
}

type customerService struct{
	CustomerRepository repository.CustomerRepository
}

//NewCustomerService membuat instansi baru dari AuthService
func NewAuthService(customerRep repository.CustomerRepository) CustomerService{
	return &customerService{
		CustomerRepository: customerRep,
	}
}

func (service *customerService) RegisterCustomer(customer dto.RegisterCustomerDTO) structs.Customer{
	customerToCreate := structs.Customer{}
	err := smapping.FillStruct(&customerToCreate, smapping.MapFields(&customer))
	if err != nil{
		log.Fatalf("Failed map %v",err)
	}

	res := service.CustomerRepository.InsertCustomer(customerToCreate)
	return res
}

func (service *customerService) IsDuplicate(customer_phone string) bool{
	res := service.CustomerRepository.IsDuplicate(customer_phone)
	return !(res.Error == nil)
}


