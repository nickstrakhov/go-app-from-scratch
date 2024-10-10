package models

type Users struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username"`
	FullName string `json:"fullname"`
}

type Prayers struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	User    Users
	Title   string `json:"title"`
	Content string `json:"content"`
}
