/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"database/sql"
	"fmt"
	"freecons/db"
	"freecons/utils"
	"log"
	checkurl "net/url"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/spf13/cobra"
)

var csvPath string
var iconURL string
var wg sync.WaitGroup

// checkupCmd represents the checkup command
var checkupCmd = &cobra.Command{
	Use:   "checkup",
	Short: "checkup incons availability online",
	Long: `Check icons availability online and update the database
	date get from by default from checkup.csv file `,
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
			wg.Add(1)
			go checkIconsAvabilityAndUpdate(database, iconURL)
			if counter == batchSize {
				wg.Wait()
				fmt.Println(time.Since(now))
				counter = 0
			}
			counter++
		}
	},
}

func init() {
	rootCmd.AddCommand(checkupCmd)
}

func checkIconsAvabilityAndUpdate(db *sql.DB, iconURL string) {
	defer wg.Done()
	statusCode, err := utils.CheckStatusCode(iconURL)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = db.Exec("UPDATE icons SET status = ? WHERE icon_url = ?", statusCode, iconURL)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No rows were updated")
		}
		log.Println(err)
		return
	}
}
