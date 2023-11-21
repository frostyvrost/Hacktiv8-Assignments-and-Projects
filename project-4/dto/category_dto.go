package dto

import (
	"net/http"
	"project-4/models"
	"project-4/pkg"
	"project-4/service"

	"github.com/gin-gonic/gin"
)

// CreateCategory godoc
// @Summary Creates a new category
// @Tags Categories
// @Accept json
// @Produce json
// @Param models.Category body models.CategoryCreate true "Category object to be created"
// @Success 201 {object} models.User "Category created successfully"
// @Failure 400 {object} pkg.ErrorResponse "Bad Request"
// @Failure 401 {object} pkg.ErrorResponse "Unauthorized"
// @Failure 422 {object} pkg.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} pkg.ErrorResponse "Server Error"
// @Router /categories [post]
func CreateCategory(context *gin.Context) {
	var category models.Category

	if err := context.ShouldBindJSON(&category); err != nil {
		errorHandler := pkg.UnprocessibleEntity("Invalid JSON body")
		context.AbortWithStatusJSON(errorHandler.Status(), errorHandler)
		return
	}

	categoryResponse, err := service.CategoryService.CreateCategory(&category)

	if err != nil {
		context.AbortWithStatusJSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"id":                  categoryResponse.ID,
		"type":                categoryResponse.Type,
		"sold_product_amount": categoryResponse.SoldProductAmount,
		"created_at":          categoryResponse.CreatedAt,
	})
}

// GetAllCategories godoc
// @Summary Get All Categories.
// @Tags Categories
// @Accept json
// @Produce json
// @Success 200 {array} models.Category "Categories fetched successfully"
// @Failure 400 {object} pkg.ErrorResponse "Bad Request"
// @Failure 401 {object} pkg.ErrorResponse "Unauthorized"
// @Failure 422 {object} pkg.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} pkg.ErrorResponse "Server Error"
// @Router /categories [get]
func GetAllCategories(context *gin.Context) {
	categories, err := service.CategoryService.GetAllCategories()

	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusOK, categories)
}

// UpdateCategory godoc
// @Summary Update a Category.
// @Tags Categories
// @Accept json
// @Produce json
// @Param categoryId path int true "Category ID"
// @Param models.Category body models.CategoryUpdate true "Category object to be updated"
// @Success 200 {object} models.User "Category updated successfully"
// @Failure 400 {object} pkg.ErrorResponse "Bad Request"
// @Failure 401 {object} pkg.ErrorResponse "Unauthorized"
// @Failure 422 {object} pkg.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} pkg.ErrorResponse "Server Error"
// @Router /categories/{categoryId} [put]
func UpdateCategory(context *gin.Context) {
	var categoryUpdated models.CategoryUpdate

	if err := context.ShouldBindJSON(&categoryUpdated); err != nil {
		errorHandler := pkg.UnprocessibleEntity("Invalid JSON body")

		context.AbortWithStatusJSON(errorHandler.Status(), errorHandler)
		return
	}

	id, _ := pkg.GetIdParam(context, "categoryId")

	categoryResponse, err := service.CategoryService.UpdateCategory(&categoryUpdated, id)

	if err != nil {
		context.AbortWithStatusJSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"id":                  categoryResponse.ID,
		"type":                categoryResponse.Type,
		"sold_product_amount": categoryResponse.SoldProductAmount,
		"updated_at":          categoryResponse.UpdatedAt,
	})
}

// DeleteCategory godoc
// @Summary Delete a Category.
// @Tags Categories
// @Accept json
// @Produce json
// @Param categoryId path int true "Category ID"
// @Success 200 {object} models.Category "Category deleted successfully"
// @Failure 400 {object} pkg.ErrorResponse "Bad Request"
// @Failure 401 {object} pkg.ErrorResponse "Unauthorized"
// @Failure 422 {object} pkg.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} pkg.ErrorResponse "Server Error"
// @Router /categories/{categoryId} [delete]
func DeleteCategory(context *gin.Context) {
	categoryId, _ := pkg.GetIdParam(context, "categoryId")

	err := service.CategoryService.DeleteCategory(categoryId)

	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Category has been successfully deleted",
	})
}
