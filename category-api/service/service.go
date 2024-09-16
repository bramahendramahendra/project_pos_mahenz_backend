package service

import (
	"fmt"
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
	// Check for duplicate category name
	isDuplicate, err := repository.CheckDuplicateCategoryName(*category.Category)
	if err != nil {
		return err
	}
	if isDuplicate {
		return fmt.Errorf("category name already exists")
	}

	return repository.CreateCategory(category)
}

func UpdateCategory(category *model.Category) error {
	// Check for duplicate category name during update
	isDuplicate, err := repository.CheckDuplicateCategoryNameForUpdate(*category.Category, category.ID)
	if err != nil {
		return err
	}
	if isDuplicate {
		return fmt.Errorf("category name already exists")
	}

	return repository.UpdateCategory(category)
}

func DeleteCategory(id uint64) error {
	return repository.DeleteCategory(id)
}

func DeletePermanentlyCategory(id uint64) error {
	return repository.DeletePermanentlyCategory(id)
}

func GetAllCategoriesWithDeleted() ([]model.Category, error) {
	return repository.GetAllCategoriesWithDeleted()
}
