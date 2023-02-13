package models

type Icons struct {
	ID      int    `json:"id"`
	URL     string `json:"url"`
	Name    string `json:"name"`
	IconURL string `json:"icon_url"`
	Status  int `json:"status"`
}
