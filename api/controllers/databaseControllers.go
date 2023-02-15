package controllers

import (
	"database/sql"
	"freecons/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

type DatabaseControllers struct {
	Database *sql.DB
}

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

func (db *DatabaseControllers) GetDashboardData(c echo.Context) error {
	var dashboard models.Dashboard
	result, err := db.Database.Query("SELECT COUNT(*) FROM icons")
	if err != nil {
		return c.JSON(500, err)
	}
	defer result.Close()
	err = result.Scan(&dashboard.TotalIcons)
	if err != nil {
		return c.JSON(500, err)
	}

	result, err = db.Database.Query("SELECT COUNT(*) FROM icons WHERE status = 200")
	if err != nil {
		return c.JSON(500, err)
	}
	defer result.Close()
	err = result.Scan(&dashboard.TotalActiveIcons)
	if err != nil {
		return c.JSON(500, err)
	}
	result, err = db.Database.Query("SELECT COUNT(*) FROM error_links")
	if err != nil {
		return c.JSON(500, err)
	}
	defer result.Close()
	err = result.Scan(&dashboard.TotalErrorIcons)
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, dashboard)
}
