package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Content string `json:"content"`
	UserID  uint   `json:"userID"`
	Status  bool   `json:"status" gorm:"default:false"`
}
