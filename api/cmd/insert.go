/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"database/sql"
	"fmt"
	"freecons/db"
	"freecons/models"
	"freecons/utils"
	"log"
	checkurl "net/url"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

// insertCmd represents the insert command
var insertCmd = &cobra.Command{
	Use:   "insert",
	Short: "insert a csv file data to database",
	Long: `insert a csv file data to database.
	you have not path csv file path args then
	insert data from our defult csv file`,
	Args: cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			csvPath = args[0]
		} else {
			csvPath = "data/checkup.csv"
		}
		rows := utils.ReadWebLinks(csvPath)
		index, err := utils.CheckWebURLColIndexFromCSV(rows)
		if err != nil {
			log.Fatalln(err)
			return
		}
		if !utils.ValidateFirst100WebURLs(rows, index) {
			log.Fatalln("your urls not valid")
			return
		}
		fmt.Println("CSV File Validated SUCCESSFULLY")

		counter := 1
		batchSize, err := strconv.Atoi(os.ExpandEnv("$CHECKUP_BATCH_SIZE"))
		if err != nil {
			log.Fatalln(err)
			return
		}
		database, err := db.Database()
		if err != nil {
			log.Fatal("Database Not Connected Due To: ", err)
		}
		if databaseEmtpy(database) {
			log.Fatalln("Database Not Empty")
			return
		}
		lastInsertdURL := getDatabaseLastInserdRow(database)
		fmt.Println("Started Inserting Data")
		defer database.Close()
		for _, row := range rows[1:] {
			now := time.Now()
			if row[index] == lastInsertdURL {
				url, err := checkurl.Parse(row[index])
				if err != nil {
					log.Fatalln(err)
					continue
				}
				if url.Scheme == "" {
					iconURL = fmt.Sprintf("https://%s/favicon.ico", url)
				} else {
					iconURL = fmt.Sprintf("%s/favicon.ico", url)
				}
				name, err := utils.GetNameServerRootFromURL(url.String())
				if err != nil {
					log.Println(err)
				}
				icons := models.Icons{
					URL:     url.String(),
					Name:    name,
					IconURL: iconURL,
				}
				insertCSVData(database, icons)
				if counter == batchSize {
					time.Sleep(2 * time.Second)
					fmt.Println(time.Since(now))
					fmt.Println(batchSize, ": Record Inserted")
					counter = 0
				}
				counter++
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(insertCmd)
}

func insertCSVData(database *sql.DB, icons models.Icons) {
	statusCode, err := utils.CheckStatusCode(icons.IconURL)
	if err != nil {
		fmt.Println(err)
		_, err = database.Exec("INSERT INTO errorlinks (icon_url) VALUES (?)", icons.IconURL)
		if err != nil {
			log.Println(err)
		}
		return
	}
	time.Sleep(500 * time.Millisecond)
	icons.Status = statusCode
	_, err = database.Exec("INSERT INTO icons (url, name, status, icon_url) VALUES (?, ?, ?, ?)", icons.URL, icons.Name, icons.Status, icons.IconURL)
	if err != nil {
		log.Println(err)
	}
}

func databaseEmtpy(database *sql.DB) bool {
	var count int
	err := database.QueryRow("SELECT COUNT(*) FROM icons").Scan(&count)
	if err != nil {
		log.Println(err)
		return false
	}
	if count == 0 {
		return true
	}
	return false
}

func getDatabaseLastInserdRow(database *sql.DB) string {
	var lastInsertdURL string
	err := database.QueryRow("SELECT url FROM icons ORDER BY id DESC LIMIT 1").Scan(&lastInsertdURL)
	if err != nil {
		log.Println(err)
		return lastInsertdURL
	}
	return lastInsertdURL
}
