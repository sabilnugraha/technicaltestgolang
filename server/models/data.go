package models

type Data struct {
	ID    int    `json:"id" gorm:"PRIMARY_KEY"`
	Image string `json:"image"  form:"image" gorm:"type: varchar(255)"`
}
