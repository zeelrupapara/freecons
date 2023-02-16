package routes

import (
	"fmt"
	"freecons/controllers"
	"freecons/db"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(app *echo.Echo) {
	database, err := db.Database()
	if err != nil {
		fmt.Println(err)
	}

	// Databse Controllers
	NewDatabaseControllers := controllers.DatabaseControllers{Database: database}
	app.GET("api/v1/icons", NewDatabaseControllers.GetIconsData)
	app.GET("api/v1/errorlinks", NewDatabaseControllers.GetErrorLinksData)
	app.GET("api/v1/dashboard/counts", NewDatabaseControllers.GetCountOfIcons)
}
