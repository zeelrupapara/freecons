/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/go-ping/ping"
	"github.com/spf13/cobra"
)

// backupCmd represents the backup command
var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		pinger, err := ping.NewPinger("google.com")
		if err != nil {
			panic(err)
		}
		pinger.Count = 3
		err = pinger.Run() // Blocks until finished.
		if err != nil {
			panic(err)
		}
		stats := pinger.Statistics()
		fmt.Println(stats)
	},
}

func init() {
	rootCmd.AddCommand(backupCmd)
}

// func CreateBackup(database *sql.DB, csvPath string) error {
// 	var tables []string
// 	var columns []string
// 	var table string
// 	var column string
// 	// Create a new CSV file
// 	file, err := os.Create(csvPath)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer file.Close()
// 	result, err := database.Query("SELECT name FROM (SELECT * FROM sqlite_schema UNION ALL SELECT * FROM sqlite_temp_schema) WHERE type='table' and name is not 'gorp_migrations' and name is not 'sqlite_sequence';")
// 	if err != nil {
// 		log.Errorf(err.Error())
// 	}
// 	for result.Next() {
// 		err = result.Scan(&table)
// 		if err != nil {
// 			log.Errorf(err.Error())
// 		}
// 		tables = append(tables, table)
// 	}
// 	for _, tableName := range tables {
// 		tableColQuery := fmt.Sprintf("select name from PRAGMA_TABLE_INFO(%s);", tableName)
// 		result, err = database.Query(tableColQuery)
// 		if err != nil {
// 			log.Errorf(err.Error())
// 		}
// 		for result.Next() {
// 			err = result.Scan(&column)
// 			if err != nil {
// 				log.Errorf(err.Error())
// 			}
// 			columns = append(columns, column)
// 		}
// 	}
// 	// Write the header row to the file
// 	writer := csv.NewWriter(file)
// 	header := []string{"column1", "column2", "column3"}
// 	if err := writer.Write(header); err != nil {
// 		panic(err.Error())
// 	}

// 	// Loop through the rows and write the data to the file
// 	// for rows.Next() {
// 	// 	var column1 string
// 	// 	var column2 string
// 	// 	var column3 string
// 	// 	if err := rows.Scan(&column1, &column2, &column3); err != nil {
// 	// 		panic(err.Error())
// 	// 	}
// 	// 	if err := writer.Write([]string{column1, column2, column3}); err != nil {
// 	// 		panic(err.Error())
// 	// 	}
// 	// }

// 	// Flush the writer to ensure all data is written to the file
// 	writer.Flush()
// }
