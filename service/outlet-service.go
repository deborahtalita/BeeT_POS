package service

import (
	"beet_pos/dto"
	"beet_pos/entity"
	"beet_pos/repository"
	"log"

	"github.com/mashingan/smapping"
)

type OutletService interface {
	CreateOutlet(outlet dto.CreateOutlet) entity.Outlet
	IsDuplicate (outlet_name string) bool
}

type outletService struct{
	OutletRepository repository.OutletRepository
}

func NewOutletService(outletRep repository.OutletRepository) OutletService{
	return &outletService{
		OutletRepository: outletRep,
	}
}

func (service *outletService) CreateOutlet(outlet dto.CreateOutlet)entity.Outlet{
	outletToCreate := entity.Outlet{}
	err := smapping.FillStruct(&outletToCreate, smapping.MapFields(&outlet))
	if err != nil{
		log.Fatalf("Failed map %v", err)
	}

	res := service.OutletRepository.InsertOutlet(outletToCreate)
	return res
}

func (service *outletService) IsDuplicate(outlet_name string) bool{
	res := service.OutletRepository.IsDuplicate(outlet_name)
	return !(res.Error == nil)
}
