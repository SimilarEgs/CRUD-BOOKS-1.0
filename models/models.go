package models

type Books struct {
	ID       int64  `json:"id"`
	BookName string `json:"bookname"`
	Author   string `json:"author"`
	Date     int64  `json:"date"`
}
