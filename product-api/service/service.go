package service

import (
	"fmt"
	"project/product-api/model"
	"project/product-api/repository"
)

func GetAllProducts() ([]model.Product, error) {
	return repository.GetAllProducts()
}

func GetProductByID(id uint64) (*model.Product, error) {
	return repository.GetProductByID(id)
}

func GetProductsByCategoryID(idCategory uint64) ([]model.Product, error) {
	return repository.GetProductsByCategoryID(idCategory)
}

func CreateProduct(product *model.Product) error {
	isDuplicate, err := repository.CheckDuplicateProductName(*product.Product)
	if err != nil {
		return err
	}
	if isDuplicate {
		return fmt.Errorf("product name already exists")
	}

	return repository.CreateProduct(product)
}

func UpdateProduct(product *model.Product) error {
	isDuplicate, err := repository.CheckDuplicateProductNameForUpdate(*product.Product, product.ID)
	if err != nil {
		return err
	}
	if isDuplicate {
		return fmt.Errorf("product name already exists")
	}

	return repository.UpdateProduct(product)
}

func DeleteProduct(id uint64) error {
	return repository.DeleteProduct(id)
}

func DeletePermanentlyProduct(id uint64) error {
	return repository.DeletePermanentlyProduct(id)
}

func GetAllProductsWithDeleted() ([]model.Product, error) {
	return repository.GetAllProductsWithDeleted()
}
