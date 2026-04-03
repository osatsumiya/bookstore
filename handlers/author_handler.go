package handlers

import (
	"net/http"

	"bookstore/models"

	"github.com/gin-gonic/gin"
)

var Authors = make(map[int]models.Author)
var NextAuthorID = 1

func GetAuthors(c *gin.Context) {
	var authorList []models.Author
	for _, author := range Authors {
		authorList = append(authorList, author)
	}
	c.JSON(http.StatusOK, authorList)
}

func CreateAuthor(c *gin.Context) {
	var author models.Author

	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if author.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "author name is required"})
		return
	}

	author.ID = NextAuthorID
	NextAuthorID++
	Authors[author.ID] = author

	c.JSON(http.StatusCreated, author)
}
