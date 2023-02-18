package controllers

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

func (database *DatabaseControllers) GetIconsList(c echo.Context) error {
	page := c.QueryParam("page")
	if page == "" {
		page = "1"
	}
	pageNum, err := strconv.Atoi(page)
	if err != nil {
		return c.JSON(400, map[string]string{
			"error": "page must be a number",
		})
	}
	result, err := database.Database.Query(`SELECT icon_url FROM icons WHERE status = 200 LIMIT 120 OFFSET ?`, 120*(pageNum-1))
	if err != nil {
		return c.JSON(500, map[string]string{
			"error": err.Error(),
		})
	}
	var icons []string
	for result.Next() {
		var icon string
		err := result.Scan(&icon)
		if err != nil {
			return c.JSON(500, map[string]string{
				"error": err.Error(),
			})
		}
		icons = append(icons, icon)
	}
	return c.JSON(200, icons)
}
