package repository

import (
	"nexmedis_project/model"

	"gorm.io/gorm"
)

type ProductRepository interface {
	FindAll() ([]model.Product, error)
	FindByID(id uint) (model.Product, error)
	Search(keyword string) ([]model.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db}
}

func (r *productRepository) FindAll() ([]model.Product, error) {
	var products []model.Product
	err := r.db.Find(&products).Error
	return products, err
}

func (r *productRepository) FindByID(id uint) (model.Product, error) {
	var product model.Product
	err := r.db.First(&product, id).Error
	return product, err
}

func (r *productRepository) Search(keyword string) ([]model.Product, error) {
	var products []model.Product
	err := r.db.Where("name LIKE ?", "%"+keyword+"%").Find(&products).Error
	return products, err
}
