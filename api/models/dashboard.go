package models

import "time"

type Dashboard struct {
	TotalIcons             int       `json:"total_icons"`
	TotalActiveIcons       int       `json:"total_active_icons"`
	TotalErrorIcons        int       `json:"total_error_icons"`
	IconsGraph             []Graph   `json:"icons_graph"`
	ErrorIconsGraph        []Graph   `json:"error_icons_graph"`
	TotalIconsPercentGraph []Graph   `json:"total_icons_percent_graph"`
	UpdatedTime            time.Time `json:"updated_time"`
}

type Graph struct {
	Labels []string `json:"labels"`
	Data   []int    `json:"data"`
}

type CountDiffIcons struct {
	TotalIcons           int    `json:"total_icons"`
	TotalIconsTime       string `json:"total_icons_time"`
	TotalActiveIcons     int    `json:"total_active_icons"`
	TotalActiveIconsTime string `json:"total_active_icons_time"`
	TotalErrorIcons      int    `json:"total_error_icons"`
	TotalErrorIconsTime  string `json:"total_error_icons_time"`
}

type LineChart struct {
	Dates  []string `json:"dates"`
	Labels []string `json:"labels"`
	Data   [][]int  `json:"data"`
}
