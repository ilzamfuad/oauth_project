package service

import (
	"nexmedis_project/model"
	"nexmedis_project/repository"
)

type ProductService interface {
	GetAllProducts() ([]model.Product, error)
	GetProductByID(id uint) (model.Product, error)
	SearchProducts(keyword string) ([]model.Product, error)
}

type productService struct {
	productRepo repository.ProductRepository
}

func NewProductService(productRepo repository.ProductRepository) ProductService {
	return &productService{productRepo}
}

func (s *productService) GetAllProducts() ([]model.Product, error) {
	return s.productRepo.FindAll()
}

func (s *productService) GetProductByID(id uint) (model.Product, error) {
	return s.productRepo.FindByID(id)
}

func (s *productService) SearchProducts(keyword string) ([]model.Product, error) {
	return s.productRepo.Search(keyword)
}
