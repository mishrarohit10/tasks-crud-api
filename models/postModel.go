package models

type Task struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Due_date    string `json:"due_date"`
	Status      string `json:"status"`
}
