package controllers

import (
	"freecons/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (db *DatabaseControllers) GetIconsData(c echo.Context) error {
	var icons []models.Icons
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
		var icon models.Icons
		err := result.Scan(&icon.ID, &icon.URL, &icon.Name, &icon.Status, &icon.IconURL)
		if err != nil {
			return c.JSON(500, err)
		}
		icons = append(icons, icon)
	}
	return c.JSON(200, icons)
}
