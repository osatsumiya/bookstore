package handlers

import (
	"net/http"

	"bookstore/models"
	"github.com/gin-gonic/gin"
)

var Categories = make(map[int]models.Category)
var NextCategoryID = 1

func GetCategories(c *gin.Context) {
	var categoryList []models.Category
	for _, category := range Categories {
		categoryList = append(categoryList, category)
	}
	c.JSON(http.StatusOK, categoryList)
}

func CreateCategory(c *gin.Context) {
	var category models.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if category.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "category name is required"})
		return
	}

	category.ID = NextCategoryID
	NextCategoryID++
	Categories[category.ID] = category

	c.JSON(http.StatusCreated, category)
}
