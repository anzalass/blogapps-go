package controller

import (
	"echodatabase/model"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func PublishArtikel(c echo.Context) error {
	blog := model.Blog{}

	c.Bind(&blog)
	blog.ID = uuid.NewString()
	err := model.DB.Save(&blog).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success created",
		"blog":    blog,
	})

}

func GetAllBlog(c echo.Context) error {
	var blogs []model.Blog

	err := model.DB.Find(&blogs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":  blogs,
		"total": len(blogs),
	})

}

func GetBlogById(c echo.Context) error {
	id := c.Param("id")

	var blogs model.Blog
	err := model.DB.First(&blogs, "id = ?", id).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": blogs,
	})

}

func UpdateBlogById(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Invalid id",
		})
	}
	blogs := model.Blog{}
	c.Bind(&blogs)
	err := model.DB.Model(&blogs).Where("id = ?", id).Updates(map[string]interface{}{"judul_blog": blogs.JudulBlog, "konten": blogs.Konten}).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Update blogs",
	})

}

func DeleteBlogById(c echo.Context) error {
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
