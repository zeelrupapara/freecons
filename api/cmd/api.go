/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"freecons/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/cobra"
)

// apiCmd represents the api command
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "run api server",
	Long:  `Run api server for run backend of freecons.`,
	Run: func(cmd *cobra.Command, args []string) {
		runAPIServer()
	},
}

func init() {
	rootCmd.AddCommand(apiCmd)
}

// For running api server
func runAPIServer() {
	app := echo.New()

	app.Use(middleware.Logger(), middleware.CORS())

	routes.SetupRoutes(app)
	
	app.Logger.Fatal(app.Start(os.ExpandEnv(":${PORT}")))
}
