package controllers

import (
	"database/sql"
	"freecons/models"
	"freecons/utils"
	"log"
	"time"
)

type DatabaseControllers struct {
	Database *sql.DB
}

func (db *DatabaseControllers) GetCountIconsData() (models.CountDiffIcons, error) {
	var countDiffIcons models.CountDiffIcons
	var totalIconsTime time.Time
	var activatedIconsTime time.Time
	var errorIconsTime time.Time

	result := db.Database.QueryRow("SELECT COUNT(*) FROM icons")
	err := result.Scan(&countDiffIcons.TotalIcons)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No rows inside icons table")
		} else {
			return countDiffIcons, err
		}
	}
	result = db.Database.QueryRow("SELECT created_at FROM icons ORDER BY created_at DESC LIMIT 1")
	err = result.Scan(&countDiffIcons.TotalActiveIconsTime)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No rows inside icons table")
		} else {
			return countDiffIcons, err
		}
	}
	totalIconsTime, err = utils.ParseTime(countDiffIcons.TotalActiveIconsTime)
	if err != nil {
		return countDiffIcons, err
	}
	countDiffIcons.TotalIconsTime = totalIconsTime.Format("2006-01-02 15:04:05")

	result = db.Database.QueryRow("SELECT COUNT(*) FROM icons WHERE status = 200")
	err = result.Scan(&countDiffIcons.TotalActiveIcons)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No rows inside acivated icons table")
		} else {
			return countDiffIcons, err
		}
	}

	result = db.Database.QueryRow("SELECT created_at FROM icons WHERE status = 200 ORDER BY created_at DESC LIMIT 1")
	err = result.Scan(&countDiffIcons.TotalActiveIconsTime)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No rows inside activated icons table")
		} else {
			return countDiffIcons, err
		}
	}
	activatedIconsTime, err = utils.ParseTime(countDiffIcons.TotalActiveIconsTime)
	if err != nil {
		return countDiffIcons, err
	}
	countDiffIcons.TotalActiveIconsTime = activatedIconsTime.Format("2006-01-02 15:04:05")

	result = db.Database.QueryRow("SELECT COUNT(*) FROM errorlinks")
	err = result.Scan(&countDiffIcons.TotalErrorIcons)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No rows inside errors icons table")
		} else {
			return countDiffIcons, err
		}
	}

	result = db.Database.QueryRow("SELECT created_at FROM errorlinks ORDER BY created_at DESC LIMIT 1")
	err = result.Scan(&countDiffIcons.TotalErrorIconsTime)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No rows inside errors icons table")
		} else {
			return countDiffIcons, err
		}
	}
	errorIconsTime, err = utils.ParseTime(countDiffIcons.TotalErrorIconsTime)
	if err != nil {
		return countDiffIcons, err
	}
	countDiffIcons.TotalErrorIconsTime = errorIconsTime.Format("2006-01-02 15:04:05")

	return countDiffIcons, nil
}
