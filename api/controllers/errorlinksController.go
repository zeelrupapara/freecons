package controllers

import (
	"freecons/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (db *DatabaseControllers) GetErrorLinksData(c echo.Context) error {
	var errorlinks []models.ErrorLinks
	page := c.QueryParam("page")
	limit := 25
	pageCount, err := strconv.Atoi(page)
	if err != nil {
		return c.JSON(500, err)
	}
	result, err := db.Database.Query("SELECT id, url, name, status, icon_url FROM icons LIMIT ? OFFSET ?", limit, (pageCount-1)*limit)
	if err != nil {
		return c.JSON(500, err)
	}
	defer result.Close()
	for result.Next() {
		var errorlink models.ErrorLinks
		err := result.Scan(&errorlink.ID, errorlink.IconURL)
		if err != nil {
			return c.JSON(500, err)
		}
		errorlinks = append(errorlinks, errorlink)
	}
	return c.JSON(200, errorlinks)
}
