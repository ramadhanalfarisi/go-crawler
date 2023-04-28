package model

type Manga struct {
	Title     string    `json:"title"`
	Genres    string    `json:"genres"`
	Year      string    `json:"year"`
	Volumes   string    `json:"volumes"`
	Chapters  string    `json:"chapters"`
	Themes    string    `json:"themes"`
	Authors   string    `json:"authors"`
	Statistic Statistic `json:"statistic"`
}

type Statistic struct {
	Score      string `json:"score"`
	Ranked     string `json:"ranked"`
	Popularity string `json:"popularity"`
}
