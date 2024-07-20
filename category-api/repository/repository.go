package repository

import (
	"database/sql"
	"project/category-api/model"
	"project/config"
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
	query := "INSERT INTO categoris (category) VALUES (?)"
	_, err := config.DB.Exec(query, category.Category)
	return err
}

func UpdateCategory(category *model.Category) error {
	query := "UPDATE categories SET category =?, updated_at CURRENT_TIMESTAMP(3) WHERE id = ? AND deleted_at IS NULL"
	_, err := config.DB.Exec(query, category.Category, category.ID)
	return err
}

func DeleteCategory(id uint64) error {
	query := "UPDATE categories SET deleted_at = CURRENT_TIMESTAMP(3) WHERE id = ? AND deleted_at IS NULL"
	_, err := config.DB.Exec(query, id)
	return err
}

func DeletePermanently(id uint64) error {
	query := "DELETE FROM categories WHERE id = ?"
	_, err := config.DB.Exec(query, id)
	return err
}
