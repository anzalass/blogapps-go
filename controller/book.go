package controller

import (
	"echodatabase/model"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func PublishBook(c echo.Context) error {
	books := model.Book{}

	c.Bind(&books)
	books.ID = uuid.NewString()
	err := model.DB.Save(&books).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success created",
		"books":   books,
	})

}

func GetAllBooks(c echo.Context) error {
	var books []model.Book

	err := model.DB.Find(&books).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":  books,
		"total": len(books),
	})

}

func GetBookById(c echo.Context) error {
	id := c.Param("id")

	var books model.Book
	err := model.DB.First(&books, "id = ?", id).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": books,
	})

}

func UpdateBookById(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Invalid id",
		})
	}
	books := model.Book{}
	c.Bind(&books)
	err := model.DB.Model(&books).Where("id = ?", id).Updates(map[string]interface{}{"judul": books.Judul, "penerbit": books.Penerbit}).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Update books",
	})

}

func DeleteBooksById(c echo.Context) error {
	id := c.Param("id")

	var books = model.Book{}
	err := model.DB.Delete(&books, "id = ?", id).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Delete books",
	})
}
