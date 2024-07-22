package handler

import (
	"net/http"
	"project/category-api/model"
	"project/category-api/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetAllCategories
// @Summary Get all categories
// @Description Get all categories
// @Tags categories
// @Accept json
// @Product json
// @Success 200 {object} model.Response
// @Router /categories [get]
func GetAllCategories(c *gin.Context) {
	categories, err := service.GetAllCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			ResponseCode: http.StatusInternalServerError,
			ResponseDesc: "Failed to fetch categories",
			ResponseData: nil,
		})
		return
	}

	response := model.Response{
		ResponseCode: http.StatusOK,
		ResponseDesc: "Request was successful",
		ResponseData: map[string]interface{}{
			"items": categories,
		},
		ResponseMeta: &model.ResponseMeta{
			Page:         1,
			PerPage:      10,
			TotalPages:   1,
			TotalRecords: len(categories),
		},
	}

	c.JSON(http.StatusOK, response)
}

// GetCategoryByID
// @Summary Get category by ID
// @Description Get category by ID
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} model.Response
// @Router /categories/{id} [get]
func GetCategoryByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			ResponseCode: http.StatusBadRequest,
			ResponseDesc: "Invalid ID",
			ResponseData: nil,
		})
		return
	}

	category, err := service.GetCategoryByID(uint64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			ResponseCode: http.StatusInternalServerError,
			ResponseDesc: "Failed to fetch category",
			ResponseData: nil,
		})
		return
	}

	if category == nil {
		c.JSON(http.StatusNotFound, model.Response{
			ResponseCode: http.StatusNotFound,
			ResponseDesc: "Category not found",
			ResponseData: nil,
		})
		return
	}

	response := model.Response{
		ResponseCode: http.StatusOK,
		ResponseDesc: "Request was successful",
		ResponseData: category,
	}

	c.JSON(http.StatusOK, response)
}

// CreateCategory
// @Summary Create a new category
// @Description Create a new category
// @Tags categories
// @Accept json
// @Produce json
// @Param category body model.CategoryInput true "Category"
// @Success 201 {object} model.Response
// @Router /categories [post]
func CreateCategory(c *gin.Context) {
	var category model.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			ResponseCode: http.StatusBadRequest,
			ResponseDesc: err.Error(),
			ResponseData: nil,
		})
		return
	}

	if err := service.CreateCategory(&category); err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			ResponseCode: http.StatusInternalServerError,
			ResponseDesc: "Failed to create category",
			ResponseData: nil,
		})
		return
	}

	response := model.Response{
		ResponseCode: http.StatusCreated,
		ResponseDesc: "Category created successfully",
		ResponseData: category,
	}

	c.JSON(http.StatusCreated, response)
}

// UpdateCategory
// @Summary Update a category
// @Description Update a category
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param category body model.Category true "Category"
// @Success 200 {object} model.Response
// @Router /categories/{id} [put]
func UpdateCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			ResponseCode: http.StatusBadRequest,
			ResponseDesc: "Invalid ID",
			ResponseData: nil,
		})
		return
	}

	var category model.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			ResponseCode: http.StatusBadRequest,
			ResponseDesc: err.Error(),
			ResponseData: nil,
		})
		return
	}

	category.ID = uint64(id)

	if err := service.UpdateCategory(&category); err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			ResponseCode: http.StatusInternalServerError,
			ResponseDesc: "Failed to update category",
			ResponseData: nil,
		})
		return
	}

	response := model.Response{
		ResponseCode: http.StatusOK,
		ResponseDesc: "Category Update successfully",
		ResponseData: category,
	}

	c.JSON(http.StatusOK, response)
}

// DeleteCategory
// @Summary Delete a category
// @Description Delete a category
// @Tags categories
// @Accept json
// @Produce json
// @Param id  path int true "Category ID"
// @Success 200 {object} model.Response
// @Router /categories/{id} [delete]
func DeleteCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			ResponseCode: http.StatusBadRequest,
			ResponseDesc: "Invalid ID",
			ResponseData: nil,
		})
		return
	}

	if err := service.DeleteCategory(uint64(id)); err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			ResponseCode: http.StatusInternalServerError,
			ResponseDesc: "Failed to delete category",
			ResponseData: nil,
		})
		return
	}

	response := model.Response{
		ResponseCode: http.StatusOK,
		ResponseDesc: "Category deleted successfully",
		ResponseData: nil,
	}
	c.JSON(http.StatusOK, response)
}

// DeletePermanentlyCategory
// @Summary Permanently delete a category
// @Description Permanently delete a category
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} model.Response
// @Router /categories/permanently/{id} [delete]
func DeletePermanentlyCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			ResponseCode: http.StatusBadRequest,
			ResponseDesc: "Invalid ID",
			ResponseData: nil,
		})
		return
	}

	if err := service.DeletePermanentlyCategory(uint64(id)); err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			ResponseCode: http.StatusInternalServerError,
			ResponseDesc: "Failed to permanently delete category",
			ResponseData: nil,
		})
		return
	}

	response := model.Response{
		ResponseCode: http.StatusOK,
		ResponseDesc: "Category permanently deleted successfully",
		ResponseData: nil,
	}

	c.JSON(http.StatusOK, response)
}
