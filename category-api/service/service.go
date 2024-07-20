package service

import (
	"project/category-api/model"
	"project/category-api/repository"
)

func GetAllCategories() ([]model.Category, error) {
	return repository.GetAllCategories()
}

func GetCategoryByID(id uint64) (*model.Category, error) {
	return repository.GetCategoryByID(id)
}
func CreateCategory(category *model.Category) error {
	return repository.CreateCategory(category)
}

func UpdateCategory(category *model.Category) error {
	return repository.UpdateCategory(category)
}

func DeleteCategory(id uint64) error {
	return repository.DeleteCategory(id)
}

func DeletePermanentlyCategory(id uint64) error {
	return repository.DeletePermanentlyCategory(id)
}
