package service

import (
	"beet_pos/dto"
	"beet_pos/entity"
	"beet_pos/repository"
	"log"

	"github.com/mashingan/smapping"
)

type ProductService interface {
	AddProduct(product dto.AddProductDTO) entity.Product
	Update(id string, product dto.UpdateProductDTO) entity.Product
	Delete(id string)
	GetAll() ([]entity.Product, error)
	GetAllPaginate(p dto.Pagination) dto.Pagination
}

type productService struct {
	productRepository repository.ProductRepository
}

func NewProductService(productRepo repository.ProductRepository) ProductService {
	return &productService{
		productRepository: productRepo,
	}
}

func (svc *productService) AddProduct(product dto.AddProductDTO) entity.Product {
	productToCreate := entity.Product{}

	err := smapping.FillStruct(&productToCreate, smapping.MapFields(&product))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := svc.productRepository.AddProduct(productToCreate)
	return res
}

func (svc *productService) Update(id string, product dto.UpdateProductDTO) entity.Product {
	productToUpdate := entity.Product{}
	err := smapping.FillStruct(&productToUpdate, smapping.MapFields(&product))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := svc.productRepository.Update(id, productToUpdate)
	return res
}

func (svc *productService) Delete(id string) {
	svc.productRepository.Delete(id)
}

// GetAll implements ProductService
func (svc *productService) GetAll() ([]entity.Product, error) {
	return svc.productRepository.GetAll()
}

// GetAllPaginate implements ProductService
func (svc *productService) GetAllPaginate(p dto.Pagination) dto.Pagination {
	return svc.productRepository.GetAllPaginate(p)
}
