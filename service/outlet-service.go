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
	UpdateOutlet(outlet dto.UpdateOutlet) entity.Outlet
	ReadOutlet() [] entity.Outlet
	FindByID(outlet_id string) entity.Outlet
	DeleteOutlet(outlet_id string) error
	IsDuplicate(outlet_name string) bool
	GetAllPaginate(p dto.Pagination) dto.Pagination
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

func (service *outletService) UpdateOutlet(outlet dto.UpdateOutlet)entity.Outlet{
	outletToUpdate := entity.Outlet{}
	err := smapping.FillStruct(&outletToUpdate, smapping.MapFields(&outlet))
	if err != nil{
		log.Fatalf("Failed map %v", err)
	}

	res := service.OutletRepository.UpdateOutlet(outletToUpdate)
	return res
}

func (service *outletService) IsDuplicate(outlet_name string) bool{
	res := service.OutletRepository.IsDuplicate(outlet_name)
	return !(res.Error == nil)
}

func (service *outletService) ReadOutlet() []entity.Outlet{
	return service.OutletRepository.ReadOutlet()
}

func (service *outletService) FindByID(outlet_id string) entity.Outlet{
	return service.OutletRepository.FindByID(outlet_id)
}

func (service *outletService) DeleteOutlet(outlet_id string) error{
	outlet := service.OutletRepository.FindByID(outlet_id)

	// if outlet == (entity.Outlet{}) {
	// 	return errors.New("Outlet tidak ada")
	// }

	err := service.OutletRepository.DeleteOutlet(outlet)
	if err != nil{
		return err
	}

	return nil
}

func (service *outletService) GetAllPaginate(p dto.Pagination) dto.Pagination {
	return service.OutletRepository.GetAllPaginate(p)
}
