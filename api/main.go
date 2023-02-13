/*
Copyright © 2023 jeelrupapara
*/
package main

import (
	"freecons/cmd"
	"freecons/utils"
)

func init() {
	// Load .env file
	utils.LoadENV()
}

func main() {
	cmd.Execute()
}
