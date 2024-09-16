package repository

import (
	"database/sql"
	"project/category-api/model"
	"project/config"
	"time"
)

func GetAllCategories() ([]model.Category, error) {
	query := "SELECT id, category, created_at, updated_at FROM categories WHERE deleted_at IS NULL"
	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := []model.Category{}
	for rows.Next() {
		var category model.Category
		if err := rows.Scan(&category.ID, &category.Category, &category.CreatedAt, &category.UpdateAt); err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}
	return categories, nil
}

func GetCategoryByID(id uint64) (*model.Category, error) {
	query := "SELECT id, category, created_at, updated_at FROM categories WHERE id = ? AND deleted_at IS NULL"
	row := config.DB.QueryRow(query, id)

	var category model.Category
	if err := row.Scan(&category.ID, &category.Category, &category.CreatedAt, &category.UpdateAt); err != nil {
		if err != sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &category, nil
}

func CreateCategory(category *model.Category) error {
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return err
	}

	now := time.Now().In(loc)
	formattedTime := now.Format("2006-01-02 15:04:05.000")
	category.CreatedAt = &formattedTime
	category.UpdateAt = &formattedTime

	query := "INSERT INTO categories (category, created_at, updated_at) VALUES (?, ?, ?)"
	result, err := config.DB.Exec(query, category.Category, category.CreatedAt, category.UpdateAt)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	category.ID = uint64(id)

	return nil
}

func UpdateCategory(category *model.Category) error {
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return err
	}

	now := time.Now().In(loc)
	formattedTime := now.Format("2006-01-02 15:04:05.000")
	category.UpdateAt = &formattedTime

	query := "UPDATE categories SET category = ?, updated_at = ? WHERE id = ? AND deleted_at IS NULL"
	_, err = config.DB.Exec(query, category.Category, category.UpdateAt, category.ID)
	if err != nil {
		return err
	}

	return nil
}

func DeleteCategory(id uint64) error {
	query := "UPDATE categories SET deleted_at = CURRENT_TIMESTAMP(3) WHERE id = ? AND deleted_at IS NULL"
	_, err := config.DB.Exec(query, id)
	return err
}

func DeletePermanentlyCategory(id uint64) error {
	query := "DELETE FROM categories WHERE id = ?"
	_, err := config.DB.Exec(query, id)
	return err
}

// CheckDuplicateCategoryName checks if a category name already exists.
func CheckDuplicateCategoryName(categoryName string) (bool, error) {
	query := "SELECT COUNT(1) FROM categories WHERE category = ? AND deleted_at IS NULL"
	var count int
	err := config.DB.QueryRow(query, categoryName).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// CheckDuplicateCategoryNameForUpdate checks if a category name already exists, excluding the current category.
func CheckDuplicateCategoryNameForUpdate(categoryName string, categoryID uint64) (bool, error) {
	query := "SELECT COUNT(1) FROM categories WHERE category = ? AND id != ? AND deleted_at IS NULL"
	var count int
	err := config.DB.QueryRow(query, categoryName, categoryID).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func GetAllCategoriesWithDeleted() ([]model.Category, error) {
	query := "SELECT id, category, created_at, updated_at, deleted_at FROM categories"
	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := []model.Category{}
	for rows.Next() {
		var category model.Category
		if err := rows.Scan(&category.ID, &category.Category, &category.CreatedAt, &category.UpdateAt, &category.DeletedAt); err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}
	return categories, nil
}
