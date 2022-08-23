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
	DeleteOutlet(outlet_id string)
	GetPaginateFiltering(p dto.Pagination) dto.Pagination
	// GetPaginateFiltering(context *gin.Context,p dto.Pagination) dto.Response
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

func (service *outletService) DeleteOutlet(outlet_id string){
	service.OutletRepository.DeleteOutlet(outlet_id)
}

func (service *outletService) GetAllPaginate(p dto.Pagination) dto.Pagination {

	return service.OutletRepository.GetAllPaginate(p)
}

// func (service *outletService) SearchOutlets(p dto.Pagination) dto.Pagination {

// 	// search query params
// 	// searchQueryParams := ""

// 	// for _, search := range p.Searchs {
// 	// 	searchQueryParams += fmt.Sprintf("&%s.%s=%s", search.Column, search.Action, search.Query)
// 	// }
// 	return service.OutletRepository.GetAllPaginate(p)
// }

func (service *outletService) GetPaginateFiltering(p dto.Pagination) dto.Pagination {

	return service.OutletRepository.GetPaginateFiltering(p)
}

// //================================== GET PAGINATE ORIGINAL ==============================================
// func (service *outletService) GetPaginateFiltering(context *gin.Context, p dto.Pagination) dto.Response{
// 	operationResult, totalPages :=  service.OutletRepository.GetPaginateFiltering(p)

// 	// if operationResult.Error != nil {
// 	// 	return dto.Response{Success: false, Message: operationResult.Error.Error()}
// 	// }

// 	var data = operationResult.Result.(dto.Pagination)

// 	// var data = operationResult.Result.(*dtos.Pagination)
// 	urlPath := context.Request.URL.Path
// 	searchQueryParams := ""

// 	for _, search := range p.Searchs {
// 		searchQueryParams += fmt.Sprintf("&%s.%s=%s", search.Column, search.Action, search.Query)
// 	}

// 	data.FirstPage = fmt.Sprintf("%s?limit=%d&page=%v&sort=%s", urlPath, p.Limit, 0, p.Sort) + searchQueryParams
// 	data.LastPage = fmt.Sprintf("%s?limit=%d&page=%v&sort=%s", urlPath, p.Limit, totalPages, p.Sort) + searchQueryParams

	
// 	if data.Page > 0 {
// 		// set previous page pagination response
// 		data.PreviousPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, p.Limit, data.Page-1, p.Sort) + searchQueryParams
// 	}

// 	if data.Page < totalPages {
// 		// set next page pagination response
// 		data.NextPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, p.Limit, data.Page+1, p.Sort) + searchQueryParams
// 	}

// 	if data.Page > totalPages {
// 		// reset previous page
// 		data.PreviousPage = ""
// 	}

// 	return dto.Response{Success: true, Data: data}
// }
