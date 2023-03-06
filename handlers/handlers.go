package handlers

import (
	"intern-assignment/database"
	"intern-assignment/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateMember(c echo.Context) error {
	member := models.Member{}

	if err := c.Bind(&member); err != nil {
		return err
	}

	database.DB.Create(&member)

	return c.JSON(http.StatusCreated, member)
}

func GetMemberList(c echo.Context) error {
	members := []models.Member{}

	if err := database.DB.Find(&members).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, members)
}

func GetMemberDetailInfo(c echo.Context) error {
	member := models.Member{}
	id := c.Param("id")

	if err := database.DB.First(&member, id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, member)
}

func UpdateMember(c echo.Context) error {
	member := models.Member{}

	if err := c.Bind(&member); err != nil {
		return err
	}

	database.DB.Save(&member)

	return c.JSON(http.StatusOK, member)
}

func DeleteMember(c echo.Context) error {
	member := models.Member{}
	id := c.Param("id")

	if err := database.DB.Delete(&member, id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.NoContent(http.StatusNoContent)
}
