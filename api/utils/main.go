package utils

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func LoadENV() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func CheckWebURLColIndexFromCSV(row [][]string) (int, error) {
	for i, col := range row[0] {
		if col == "web-urls" {
			return i, nil
		}
	}
	return 0, fmt.Errorf("web-urls column not found (please change colname to web-urls in your wensite column)")
}

func CheckStatusCode(url string) (int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	return resp.StatusCode, nil
}

func ReadWebLinks(checkUpFilePath string) [][]string {
	web_links, err := os.Open(checkUpFilePath)
	if err != nil {
		log.Fatalf("Cannot open '%s': %s\n", checkUpFilePath, err.Error())
	}
	defer web_links.Close()
	reader := csv.NewReader(bufio.NewReader(web_links))
	reader.Comma = ';'
	reader.LazyQuotes = true
	rows, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Cannot read CSV data:", err.Error())
	}
	return rows
}

func GetNameServerRootFromURL(urlString string) (string, error) {
	var name string
	u, err := url.Parse(urlString)
	if err != nil {
		fmt.Printf("Error parsing URL: %v\n", err)
		return name, err
	}
	hostname := u.Hostname()
	parts := strings.Split(hostname, ".")

	if len(parts) >= 3 {
		name = parts[1]
		return name, nil
	} else {
		name = parts[0]
		return name, nil
	}
}
