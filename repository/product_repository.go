package repository

import (
	"context"
	"repopatterngin/helper"
	"repopatterngin/model"
)

type ProductRepository interface {
	Insert(ctx context.Context, product model.Product) (model.Product, error)
	FindAll(ctx context.Context) []model.Product
	FindId(ctx context.Context, id int, product model.Product) (model.Product, error)
	Update(ctx context.Context, id int, product model.Product, productUpdate model.ProductUpdate) (model.Product, error)
	Delete(ctx context.Context, id int) (model.Product, error)
}

type ProductsRepository struct {
}

func StartProductRepository() ProductRepository {
	return &ProductsRepository{}
}

func (repository *ProductsRepository) Insert(ctx context.Context, product model.Product) (model.Product, error) {
	tx := helper.DB.WithContext(ctx)
	tx.Create(&product)

	return product, nil
}

func (repository *ProductsRepository) FindAll(ctx context.Context) []model.Product {
	var product []model.Product
	tx := helper.DB.WithContext(ctx)
	tx.Find(&product)

	return product
}

func (repository *ProductsRepository) FindId(ctx context.Context, id int, product model.Product) (model.Product, error) {
	product.Id = id
	tx := helper.DB.WithContext(ctx)
	tx.First(&product, id)

	return product, nil
}

func (repository *ProductsRepository) Update(ctx context.Context, id int, product model.Product, productUpdate model.ProductUpdate) (model.Product, error) {
	product.Id = id

	tx := helper.DB.WithContext(ctx)
	tx.Model(&product).Updates(productUpdate.ToModel())

	return product, nil
}

func (repository *ProductsRepository) Delete(ctx context.Context, id int) (model.Product, error) {
	var product model.Product
	product.Id = id

	tx := helper.DB.WithContext(ctx)
	tx.First(&product, id)
	tx.Delete(&product)

	return product, nil
}
