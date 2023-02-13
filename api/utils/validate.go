package utils

import (
	"fmt"
	urlcheck "net/url"
)

func ValidateURL(u string) bool {
	url, err := urlcheck.Parse(u)
	if err != nil {
		return false
	}

	if url.Scheme == "" {
		u = fmt.Sprintf("https://%s", u)
	}

	_, err = urlcheck.ParseRequestURI(u)
	if err != nil {
		fmt.Println()
		return false
	}
	return true
}

func ValidateFirst100WebURLs(row [][]string, index int) bool {
	for i := 1; i < 100; i++ {
		if !ValidateURL(row[i][index]) {
			fmt.Println(i, row[i][index], false)
			return false
		}
	}
	return true
}
