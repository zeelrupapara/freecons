package controllers

import (
	"freecons/models"
	"freecons/utils"

	"github.com/labstack/echo/v4"
)

func (db *DatabaseControllers) GetPieChart(c echo.Context) error {
	var totalIcons int
	var totalActiveIcons int
	var totalErrorIcons int
	var pieChart models.Graph
	result := db.Database.QueryRow("SELECT COUNT(*) FROM icons WHERE not status = 200")
	err := result.Scan(&totalIcons)
	if err != nil {
		return c.JSON(500, err)
	}
	result = db.Database.QueryRow("SELECT COUNT(*) FROM icons WHERE status = 200")
	err = result.Scan(&totalActiveIcons)
	if err != nil {
		return c.JSON(500, err)
	}
	result = db.Database.QueryRow("SELECT COUNT(*) FROM errorlinks")
	err = result.Scan(&totalErrorIcons)
	if err != nil {
		return c.JSON(500, err)
	}
	total := totalIcons + totalActiveIcons + totalErrorIcons
	totalIconsPer := utils.CalculatePercentage(totalIcons, total) + "%"
	totalActiveIconsPer := utils.CalculatePercentage(totalActiveIcons, total) + "%"
	totalErrorIconsPer := utils.CalculatePercentage(totalErrorIcons, total) + "%"
	pieChart.Data = []int{totalIcons, totalActiveIcons, totalErrorIcons}
	pieChart.Labels = []string{totalIconsPer, totalActiveIconsPer, totalErrorIconsPer}
	return c.JSON(200, pieChart)
}

func (db *DatabaseControllers) GetLineChart(c echo.Context) error {
	var Dates []string
	var currentDate string
	var countIconsPerDay []int
	var countActivatedIconsPerDay []int
	var countErrorIconsPerDay []int
	var countIcons int
	var countActivatedIcons int
	var countErrorIcons int

	date := c.QueryParam("date")
	// get dates
	result, err := db.Database.Query("SELECT DISTINCT DATE(created_at) FROM icons ORDER BY created_at DESC")
	if err != nil {
		return c.JSON(500, err)
	}
	defer result.Close()
	for result.Next() {
		var date string
		err := result.Scan(&date)
		if err != nil {
			return c.JSON(500, err)
		}
		Dates = append(Dates, date)
	}
	if date != "" {
		currentDate = date
	} else {
		currentDate = Dates[0]
	}
	timeStamps := utils.GetTwoHoursTimestamp()
	for i := 1; i < len(timeStamps); i++ {
		result := db.Database.QueryRow("SELECT COUNT(*) FROM icons WHERE created_at BETWEEN ? AND ?;", currentDate+" "+timeStamps[i-1], currentDate+" "+timeStamps[i])
		err := result.Scan(&countIcons)
		if err != nil {
			return c.JSON(500, err)
		}
		countIconsPerDay = append(countIconsPerDay, countIcons)

		result = db.Database.QueryRow("SELECT COUNT(*) FROM icons WHERE status = 200 AND created_at BETWEEN ? AND ?;", currentDate+" "+timeStamps[i-1], currentDate+" "+timeStamps[i])
		err = result.Scan(&countActivatedIcons)
		if err != nil {
			return c.JSON(500, err)
		}
		countActivatedIconsPerDay = append(countActivatedIconsPerDay, countActivatedIcons)
		result = db.Database.QueryRow("SELECT COUNT(*) FROM errorlinks WHERE created_at BETWEEN ? AND ?;", currentDate+" "+timeStamps[i-1], currentDate+" "+timeStamps[i])
		err = result.Scan(&countErrorIcons)
		if err != nil {
			return c.JSON(500, err)
		}
		countErrorIconsPerDay = append(countErrorIconsPerDay, countErrorIcons)
	}
	data := [][]int{countIconsPerDay, countActivatedIconsPerDay, countErrorIconsPerDay}
	timeStamps = timeStamps[1:]
	lineChart := models.LineChart{
		Dates:  Dates,
		Labels: timeStamps,
		Data:   data}
	return c.JSON(200, lineChart)
}
