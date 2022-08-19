package service

import (
	"beet_pos/dto"
	"beet_pos/entity"
	"beet_pos/repository"
	"fmt"
	"log"

	"github.com/mashingan/smapping"
)

type ProductService interface {
	AddProduct(product dto.AddProductDTO) entity.Product
	Update(id string, product dto.UpdateProductDTO) entity.Product
	Delete(id string)
	FindByID(id string) entity.Product
	GetAll() ([]entity.Product, error)
	GetAllPaginate(outlet_id string, p dto.Pagination) dto.Pagination
	AddVariant(variant dto.AddVariantDTO, id string) entity.Product_variant
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
func (svc *productService) GetAllPaginate(outlet_id string, p dto.Pagination) dto.Pagination {
	return svc.productRepository.GetAllPaginate(outlet_id, p)
}

// FindByID implements ProductService
func (svc *productService) FindByID(id string) entity.Product {
	return svc.productRepository.FindByID(id)
}

// AddVariant implements ProductService
func (svc *productService) AddVariant(variant dto.AddVariantDTO, id string) entity.Product_variant {
	variantToCreate := entity.Product_variant{}

	err := smapping.FillStruct(&variantToCreate, smapping.MapFields(&variant))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	variantToCreate.Product_id = id
	fmt.Printf("svc %d",variantToCreate.Product_id)
	res := svc.productRepository.AddVariant(variantToCreate, id)
	return res
}
