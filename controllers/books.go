package controllers

import (
	"net/http"
	"quiz3/database"
	"quiz3/repository"
	"quiz3/structs"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllBooks(c *gin.Context) {
	var (
		result gin.H
	)

	books, err := repository.GetAllBooks(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": books,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertBooks(c *gin.Context) {
	var books structs.Books

	err := c.ShouldBindJSON(&books)
	if err != nil {
		panic(err)
	}

	err = repository.InsertBooks(database.DbConnection, books)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Book",
	})

}

func UpdatedBook(c *gin.Context) {
	var books structs.Books
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&books)
	if err != nil {
		panic(err)
	}

	books.ID = int64(id)

	err = repository.UpdatedBooks(database.DbConnection, books)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update Book",
	})
}

func DeletedBook(c *gin.Context) {
	var books structs.Books
	id, err := strconv.Atoi(c.Param("id"))

	books.ID = int64(id)

	err = repository.DeletedBooks(database.DbConnection, books)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Deleted Book",
	})
}
