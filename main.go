package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book

func main() {
	r := gin.Default()

	r.GET("/books", getBooks)
	r.GET("/books/:id", getBookByID)
	r.POST("/books", createBook)
	r.PUT("/books/:id", updateBook)
	r.DELETE("/books/:id", deleteBook)

	r.Run(":8080")
}

func getBooks(c *gin.Context) {
	c.JSON(http.StatusOK, books)
}

func getBookByID(c *gin.Context) {
	id := c.Param("id")

	for _, book := range books {
		if book.ID == id {
			c.JSON(http.StatusOK, book)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}

func createBook(c *gin.Context) {
	var book Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	books = append(books, book)

	c.JSON(http.StatusCreated, book)
}

func updateBook(c *gin.Context) {
	id := c.Param("id")

	var updatedBook Book
	if err := c.ShouldBindJSON(&updatedBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, book := range books {
		if book.ID == id {
			books[i] = updatedBook
			c.JSON(http.StatusOK, gin.H{"message": "Book updated successfully"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}

func deleteBook(c *gin.Context) {
	id := c.Param("id")

	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}
