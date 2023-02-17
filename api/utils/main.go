package utils

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

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

func StructToBytes(s interface{}) []byte {
	b := new(bytes.Buffer)
	binary.Write(b, binary.LittleEndian, s)
	return b.Bytes()
}

func ParseTime(t string) (time.Time, error) {
	var pt time.Time
	pt, err := time.Parse(time.RFC3339, t)
	if err != nil {
		return pt, err
	}
	return pt, nil
}

func CalculatePercentage(count int, totalCount int) string {
	percentage := (count * 100) / totalCount
	return strconv.Itoa(percentage)
}

func GetTwoHoursTimestamp() []string {
	timeStamp := []string{"00:00:00", "02:00:00", "04:00:00", "06:00:00", "08:00:00", "10:00:00", "12:00:00", "14:00:00", "16:00:00", "18:00:00", "20:00:00", "22:00:00", "23:59:59"}
	return timeStamp
}

func EncodeString(s string) string {
	encodedStr := base64.StdEncoding.EncodeToString([]byte(s))
	return encodedStr
}

func DecodeString(s string) string {
	decodedStr, _ := base64.StdEncoding.DecodeString(s)
	return string(decodedStr)
}
