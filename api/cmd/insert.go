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
		fmt.Println("Started Inserting Data")
		batchSize, err := strconv.Atoi(os.ExpandEnv("$CHECKUP_BATCH_SIZE"))
		if err != nil {
			log.Fatalln(err)
			return
		}
		database, err := db.Database()
		if err != nil {
			log.Fatal("Database Not Connected Due To: ", err)
		}
		defer database.Close()
		for _, row := range rows[1:] {
			now := time.Now()
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
			// wg.Add(1)
			insertCSVData(database, icons)
			if counter == batchSize {
				// wg.Wait()
				time.Sleep(5 * time.Second)
				fmt.Println(time.Since(now))
				fmt.Println(batchSize, ": Record Inserted")
				counter = 0
			}
			counter++
		}

	},
}

func init() {
	rootCmd.AddCommand(insertCmd)
}

func insertCSVData(database *sql.DB, icons models.Icons) {
	// defer wg.Done()
	statusCode, err := utils.CheckStatusCode(icons.IconURL)
	if err != nil {
		fmt.Println(err)
		_, err = database.Exec("INSERT INTO errorlinks (icon_url) VALUES (?)", icons.IconURL)
		if err != nil {
			log.Println(err)
		}
		return
	}
	icons.Status = statusCode
	_, err = database.Exec("INSERT INTO icons (url, name, status, icon_url) VALUES (?, ?, ?, ?)", icons.URL, icons.Name, icons.Status, icons.IconURL)
	if err != nil {
		log.Println(err)
	}
}
