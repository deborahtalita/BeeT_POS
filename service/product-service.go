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
}

type productService struct {
	ProductRepository repository.ProductRepository
}

func NewProductService(productRepo repository.ProductRepository) ProductService {
	return &productService{
		ProductRepository: productRepo,
	}
}

func (svc *productService)AddProduct(product dto.AddProductDTO) entity.Product {
	productToCreate := entity.Product {}
	err := smapping.FillStruct(&productToCreate, smapping.MapFields(&product))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := svc.ProductRepository.AddProduct(productToCreate)
	return res
}