/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"freecons/db"
	"log"

	migrate "github.com/rubenv/sql-migrate"
	"github.com/spf13/cobra"
)

// migrateCmd represents the migrate command
func getMigrationCMD() *cobra.Command {
	migrateCmd := &cobra.Command{
		Use:   "migrate",
		Short: "migration database",
		Long:  `For appling migration in database`,
		Args:  cobra.MinimumNArgs(1),
	}
	return migrateCmd
}

func getMigrationSubCommand() []*cobra.Command {
	var migrate_sub_commands []*cobra.Command
	migrateUp := &cobra.Command{
		Use:   "up",
		Short: "It will apply migration(s)",
		Long:  `It will run all remaining migration(s)`,
		Args:  cobra.MinimumNArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			migrationUp()
			return nil
		},
	}

	migrateDown := &cobra.Command{
		Use:   "down",
		Short: "It will revert migration(s)",
		Long:  `It will run all remaining migration(s)`,
		Args:  cobra.MinimumNArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			migrationDown()
			return nil
		},
	}
	migrate_sub_commands = append(migrate_sub_commands, migrateUp, migrateDown)
	return migrate_sub_commands
}

func init() {
	migrateCmd := getMigrationCMD()
	migrate_sub_commands := getMigrationSubCommand()
	migrateCmd.AddCommand(migrate_sub_commands...)
	rootCmd.AddCommand(migrateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// migrateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// migrateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func migrationUp() {
	database, err := db.Database()
	if err != nil {
		log.Fatal("Database Not Connected Due To: ", err)
	}
	defer database.Close()
	fmt.Println("Waiting For Migrations...")
	migrations := &migrate.FileMigrationSource{
		Dir: "db/migrations",
	}
	// Apply Migration
	n, err := migrate.Exec(database, "sqlite3", migrations, migrate.Up)
	if err != nil {
		log.Fatal("Migration Not Apply Due To: ", err)
	}
	fmt.Printf("Applied %d migrations!\n", n)
}

func migrationDown() {
	database, err := db.Database()
	if err != nil {
		log.Fatal("Database Not Connected Due To: ", err)
	}
	defer database.Close()
	fmt.Println("Waiting For Migrations...")
	migrations := &migrate.FileMigrationSource{
		Dir: "db/migrations",
	}
	// Apply Migration
	n, err := migrate.Exec(database, "sqlite3", migrations, migrate.Down)
	if err != nil {
		log.Fatal("Migration Not Apply Due To: ", err)
	}
	fmt.Printf("Applied %d migrations!\n", n)
}
