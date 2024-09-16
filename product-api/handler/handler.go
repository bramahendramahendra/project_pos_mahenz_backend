package handler

import (
	"net/http"
	"project/product-api/model"
	"project/product-api/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetAllProducts
// @Summary Get all products
// @Description Get all products
// @Tags products
// @Accept json
// @Product json
// @Success 200 {object} model.Response
// @Router /products [get]
func GetAllProducts(c *gin.Context) {
	products, err := service.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			ResponseCode: http.StatusInternalServerError,
			ResponseDesc: "Failed to fetch products",
			ResponseData: nil,
		})
		return
	}

	response := model.Response{
		ResponseCode: http.StatusOK,
		ResponseDesc: "Request was successful",
		ResponseData: map[string]interface{}{
			"items": products,
		},
		ResponseMeta: &model.ResponseMeta{
			Page:         1,
			PerPage:      10,
			TotalPages:   1,
			TotalRecords: len(products),
		},
	}

	c.JSON(http.StatusOK, response)
}

// GetProductByID
// @Summary Get product by ID
// @Description Get product by ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} model.Response
// @Router /products/{id} [get]
func GetProductByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			ResponseCode: http.StatusBadRequest,
			ResponseDesc: "Invalid ID",
			ResponseData: nil,
		})
		return
	}

	product, err := service.GetProductByID(uint64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			ResponseCode: http.StatusInternalServerError,
			ResponseDesc: "Failed to fetch product",
			ResponseData: nil,
		})
		return
	}

	if product == nil {
		c.JSON(http.StatusNotFound, model.Response{
			ResponseCode: http.StatusNotFound,
			ResponseDesc: "Product not found",
			ResponseData: nil,
		})
		return
	}

	response := model.Response{
		ResponseCode: http.StatusOK,
		ResponseDesc: "Request was successful",
		ResponseData: product,
	}

	c.JSON(http.StatusOK, response)
}

// GetProductsByCategoryID
// @Summary Get products by category ID
// @Description Get products by category ID
// @Tags products
// @Accept json
// @Produce json
// @Param id_category path int true "Category ID"
// @Success 200 {object} model.Response
// @Router /products/category/{id_category} [get]
func GetProductsByCategoryID(c *gin.Context) {
	idCategory, err := strconv.Atoi(c.Param("id_category"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			ResponseCode: http.StatusBadRequest,
			ResponseDesc: "Invalid Category ID",
			ResponseData: nil,
		})
		return
	}

	products, err := service.GetProductsByCategoryID(uint64(idCategory))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			ResponseCode: http.StatusInternalServerError,
			ResponseDesc: "Failed to fetch products",
			ResponseData: nil,
		})
		return
	}

	if len(products) == 0 {
		c.JSON(http.StatusNotFound, model.Response{
			ResponseCode: http.StatusNotFound,
			ResponseDesc: "No products found for this category",
			ResponseData: nil,
		})
		return
	}

	response := model.Response{
		ResponseCode: http.StatusOK,
		ResponseDesc: "Request was successful",
		ResponseData: map[string]interface{}{
			"items": products,
		},
	}

	c.JSON(http.StatusOK, response)
}

// CreateProduct
// @Summary Create a new product
// @Description Create a new product
// @Tags products
// @Accept json
// @Produce json
// @Param product body model.ProductInput true "Product"
// @Success 201 {object} model.Response
// @Router /products [post]
func CreateProduct(c *gin.Context) {
	var product model.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			ResponseCode: http.StatusBadRequest,
			ResponseDesc: err.Error(),
			ResponseData: nil,
		})
		return
	}

	if err := service.CreateProduct(&product); err != nil {
		if err.Error() == "product name already exists" {
			c.JSON(http.StatusConflict, model.Response{
				ResponseCode: http.StatusConflict,
				ResponseDesc: "Product name already exists",
				ResponseData: nil,
			})
			return
		}

		c.JSON(http.StatusInternalServerError, model.Response{
			ResponseCode: http.StatusInternalServerError,
			ResponseDesc: "Failed to create product",
			ResponseData: nil,
		})
		return
	}

	response := model.Response{
		ResponseCode: http.StatusCreated,
		ResponseDesc: "Product created successfully",
		ResponseData: product,
	}

	c.JSON(http.StatusCreated, response)
}

// UpdateProduct
// @Summary Update a product
// @Description Update a product
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body model.ProductInput true "Product"
// @Success 200 {object} model.Response
// @Router /products/{id} [put]
func UpdateProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			ResponseCode: http.StatusBadRequest,
			ResponseDesc: "Invalid ID",
			ResponseData: nil,
		})
		return
	}

	var product model.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			ResponseCode: http.StatusBadRequest,
			ResponseDesc: err.Error(),
			ResponseData: nil,
		})
		return
	}

	existingProduct, err := service.GetProductByID(uint64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			ResponseCode: http.StatusInternalServerError,
			ResponseDesc: "Failed to fetch product",
			ResponseData: nil,
		})
		return
	}

	if existingProduct.DeletedAt != nil {
		c.JSON(http.StatusForbidden, model.Response{
			ResponseCode: http.StatusForbidden,
			ResponseDesc: "Product has been deleted and cannot be updated",
			ResponseData: nil,
		})
		return
	}

	product.ID = uint64(id)
	if err := service.UpdateProduct(&product); err != nil {
		if err.Error() == "product name already exists" {
			c.JSON(http.StatusConflict, model.Response{
				ResponseCode: http.StatusConflict,
				ResponseDesc: "Product name already exists",
				ResponseData: nil,
			})
			return
		}

		c.JSON(http.StatusInternalServerError, model.Response{
			ResponseCode: http.StatusInternalServerError,
			ResponseDesc: "Failed to update product",
			ResponseData: nil,
		})
		return
	}

	response := model.Response{
		ResponseCode: http.StatusOK,
		ResponseDesc: "Product Update successfully",
		ResponseData: product,
	}
	c.JSON(http.StatusOK, response)
}

// DeleteProduct
// @Summary Delete a product
// @Description Delete a product
// @Tags products
// @Accept json
// @Produce json
// @Param id  path int true "Product ID"
// @Success 200 {object} model.Response
// @Router /products/{id} [delete]
func DeleteProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			ResponseCode: http.StatusBadRequest,
			ResponseDesc: "Invalid ID",
			ResponseData: nil,
		})
		return
	}

	if err := service.DeleteProduct(uint64(id)); err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			ResponseCode: http.StatusInternalServerError,
			ResponseDesc: "Failed to delete product",
			ResponseData: nil,
		})
		return
	}

	response := model.Response{
		ResponseCode: http.StatusOK,
		ResponseDesc: "Product deleted successfully",
		ResponseData: nil,
	}
	c.JSON(http.StatusOK, response)
}

// DeletePermanentlyProduct
// @Summary Permanently delete a product
// @Description Permanently delete a product
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} model.Response
// @Router /products/permanently/{id} [delete]
func DeletePermanentlyProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			ResponseCode: http.StatusBadRequest,
			ResponseDesc: "Invalid ID",
			ResponseData: nil,
		})
		return
	}

	if err := service.DeletePermanentlyProduct(uint64(id)); err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			ResponseCode: http.StatusInternalServerError,
			ResponseDesc: "Failed to permanently delete product",
			ResponseData: nil,
		})
		return
	}

	response := model.Response{
		ResponseCode: http.StatusOK,
		ResponseDesc: "Product permanently deleted successfully",
		ResponseData: nil,
	}

	c.JSON(http.StatusOK, response)
}

// GetAllProductsWithDeleted
// @Summary Get all products including deleted
// @Description Get all products including deleted
// @Tags products
// @Accept json
// @Product json
// @Success 200 {object} model.Response
// @Router /products/with-deleted [get]
func GetAllProductsWithDeleted(c *gin.Context) {
	products, err := service.GetAllProductsWithDeleted()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			ResponseCode: http.StatusInternalServerError,
			ResponseDesc: "Failed to fetch products",
			ResponseData: nil,
		})
		return
	}

	response := model.Response{
		ResponseCode: http.StatusOK,
		ResponseDesc: "Request was successful",
		ResponseData: map[string]interface{}{
			"items": products,
		},
	}

	c.JSON(http.StatusOK, response)
}
