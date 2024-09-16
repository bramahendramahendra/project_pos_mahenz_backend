package repository

import (
	"database/sql"
	"project/config"
	"project/product-api/model"
	"time"
)

func GetAllProducts() ([]model.Product, error) {
	query := "SELECT id, id_category, product, created_at, updated_at FROM products WHERE deleted_at IS NULL"
	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []model.Product{}
	for rows.Next() {
		var product model.Product
		if err := rows.Scan(&product.ID, &product.IDCategory, &product.Product, &product.CreatedAt, &product.UpdatedAt); err != nil {
			return nil, err
		}

		products = append(products, product)
	}
	return products, nil
}

func GetAllProductsWithDeleted() ([]model.Product, error) {
	query := "SELECT id, id_category, product, created_at, updated_at, deleted_at FROM products"
	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []model.Product{}
	for rows.Next() {
		var product model.Product
		if err := rows.Scan(&product.ID, &product.IDCategory, &product.Product, &product.CreatedAt, &product.UpdatedAt, &product.DeletedAt); err != nil {
			return nil, err
		}

		products = append(products, product)
	}
	return products, nil
}

func GetProductByID(id uint64) (*model.Product, error) {
	query := "SELECT id, id_category, product, created_at, updated_at FROM products WHERE id = ? AND deleted_at IS NULL"
	row := config.DB.QueryRow(query, id)

	var product model.Product
	if err := row.Scan(&product.ID, &product.IDCategory, &product.Product, &product.CreatedAt, &product.UpdatedAt); err != nil {
		if err != sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &product, nil
}

func GetProductsByCategoryID(idCategory uint64) ([]model.Product, error) {
	query := "SELECT id, id_category, product, created_at, updated_at FROM products WHERE id_category = ? AND deleted_at IS NULL"
	rows, err := config.DB.Query(query, idCategory)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []model.Product{}
	for rows.Next() {
		var product model.Product
		if err := rows.Scan(&product.ID, &product.IDCategory, &product.Product, &product.CreatedAt, &product.UpdatedAt); err != nil {
			return nil, err
		}

		products = append(products, product)
	}
	return products, nil
}

func CheckDuplicateProductName(productName string) (bool, error) {
	query := "SELECT COUNT(1) FROM products WHERE product = ? AND deleted_at IS NULL"
	var count int
	err := config.DB.QueryRow(query, productName).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func CreateProduct(product *model.Product) error {
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return err
	}

	now := time.Now().In(loc)
	formattedTime := now.Format("2006-01-02 15:04:05.000")
	product.CreatedAt = &formattedTime
	product.UpdatedAt = &formattedTime

	query := "INSERT INTO products (id_category, product, created_at, updated_at) VALUES (?, ?, ?, ?)"
	result, err := config.DB.Exec(query, product.IDCategory, product.Product, product.CreatedAt, product.UpdatedAt)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	product.ID = uint64(id)

	return nil
}

func CheckDuplicateProductNameForUpdate(productName string, productID uint64) (bool, error) {
	query := "SELECT COUNT(1) FROM products WHERE product = ? AND id != ? AND deleted_at IS NULL"
	var count int
	err := config.DB.QueryRow(query, productName, productID).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func UpdateProduct(product *model.Product) error {
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return err
	}

	now := time.Now().In(loc)
	formattedTime := now.Format("2006-01-02 15:04:05.000")
	product.UpdatedAt = &formattedTime

	query := "UPDATE products SET id_category = ?, product = ?, updated_at = ? WHERE id = ? AND deleted_at IS NULL"
	_, err = config.DB.Exec(query, product.IDCategory, product.Product, product.UpdatedAt, product.ID)
	if err != nil {
		return err
	}

	return nil
}

func DeleteProduct(id uint64) error {
	query := "UPDATE products SET deleted_at = CURRENT_TIMESTAMP(3) WHERE id = ? AND deleted_at IS NULL"
	_, err := config.DB.Exec(query, id)
	return err
}

func DeletePermanentlyProduct(id uint64) error {
	query := "DELETE FROM products WHERE id = ?"
	_, err := config.DB.Exec(query, id)
	return err
}
