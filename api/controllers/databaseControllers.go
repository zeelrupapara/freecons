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

func (db *DatabaseControllers) GetCountIconsData() (models.CountDiffIcons, error) {
	var countDiffIcons models.CountDiffIcons
	result := db.Database.QueryRow("SELECT COUNT(*) FROM icons")
	err := result.Scan(&countDiffIcons.TotalIcons)
	if err != nil {
		return countDiffIcons, err
	}
	result = db.Database.QueryRow("SELECT created_at FROM icons ORDER BY created_at DESC LIMIT 1")
	err = result.Scan(&countDiffIcons.TotalIconsTime)
	if err != nil {
		return countDiffIcons, err
	}

	result = db.Database.QueryRow("SELECT COUNT(*) FROM icons WHERE status = 200")
	err = result.Scan(&countDiffIcons.TotalActiveIcons)
	if err != nil {
		return countDiffIcons, err
	}

	result = db.Database.QueryRow("SELECT created_at FROM icons WHERE status = 200 ORDER BY created_at DESC LIMIT 1")
	err = result.Scan(&countDiffIcons.TotalActiveIconsTime)
	if err != nil {
		return countDiffIcons, err
	}

	result = db.Database.QueryRow("SELECT COUNT(*) FROM errorlinks")
	err = result.Scan(&countDiffIcons.TotalErrorIcons)
	if err != nil {
		return countDiffIcons, err
	}

	result = db.Database.QueryRow("SELECT created_at FROM errorlinks ORDER BY created_at DESC LIMIT 1")
	err = result.Scan(&countDiffIcons.TotalErrorIconsTime)
	if err != nil {
		return countDiffIcons, err
	}

	return countDiffIcons, nil
}
