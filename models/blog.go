package models

type Blog struct {
	Id       int    `json:"id"`
	Avatar   string `json:"avatar"`
	Title    string `json:"title"`
	Text     string `json:"text"`
	Webpct   string `json:"webpct"`
	Url      string `json:"url"`
	Examine  string `json:"examine"`
	Category string `json:"category"`
}

func (Blog) TableName() string {
	return "blog"
}
