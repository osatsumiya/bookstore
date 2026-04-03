package handlers

import (
	"net/http"
	"strconv"

	"bookstore/models"

	"github.com/gin-gonic/gin"
)

var Books = make(map[int]models.Book)
var NextBookID = 1

func GetBooks(c *gin.Context) {
	var bookList []models.Book

	categoryParam := c.Query("category")
	authorParam := c.Query("author")
	pageParam := c.DefaultQuery("page", "1")
	limitParam := c.DefaultQuery("limit", "10")

	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitParam)
	if err != nil || limit < 1 {
		limit = 10
	}

	var categoryIDFilter int
	var authorIDFilter int

	if categoryParam != "" {
		categoryIDFilter, _ = strconv.Atoi(categoryParam)
	}

	if authorParam != "" {
		authorIDFilter, _ = strconv.Atoi(authorParam)
	}

	for _, book := range Books {
		if categoryParam != "" && book.CategoryID != categoryIDFilter {
			continue
		}
		if authorParam != "" && book.AuthorID != authorIDFilter {
			continue
		}
		bookList = append(bookList, book)
	}

	start := (page - 1) * limit
	end := start + limit

	if start > len(bookList) {
		start = len(bookList)
	}
	if end > len(bookList) {
		end = len(bookList)
	}

	c.JSON(http.StatusOK, gin.H{
		"page":  page,
		"limit": limit,
		"total": len(bookList),
		"data":  bookList[start:end],
	})
}

func GetBookByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid book id"})
		return
	}

	book, exists := Books[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
		return
	}

	c.JSON(http.StatusOK, book)
}

func CreateBook(c *gin.Context) {
	var book models.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if book.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "title is required"})
		return
	}

	if book.Price <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "price must be greater than 0"})
		return
	}

	if _, exists := Authors[book.AuthorID]; !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "author does not exist"})
		return
	}

	if _, exists := Categories[book.CategoryID]; !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "category does not exist"})
		return
	}

	book.ID = NextBookID
	NextBookID++
	Books[book.ID] = book

	c.JSON(http.StatusCreated, book)
}

func UpdateBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid book id"})
		return
	}

	_, exists := Books[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
		return
	}

	var updatedBook models.Book
	if err := c.ShouldBindJSON(&updatedBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if updatedBook.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "title is required"})
		return
	}

	if updatedBook.Price <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "price must be greater than 0"})
		return
	}

	if _, exists := Authors[updatedBook.AuthorID]; !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "author does not exist"})
		return
	}

	if _, exists := Categories[updatedBook.CategoryID]; !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "category does not exist"})
		return
	}

	updatedBook.ID = id
	Books[id] = updatedBook

	c.JSON(http.StatusOK, updatedBook)
}

func DeleteBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid book id"})
		return
	}

	_, exists := Books[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
		return
	}

	delete(Books, id)
	c.JSON(http.StatusOK, gin.H{"message": "book deleted successfully"})
}
