package models

import (
	"time"
)

type Task struct {
	Id          int `gorm:"primaryKey"`
	Code        string
	Description string
	AssignedTo  string
	Completed   bool
	CreatedBy   string
	UpdatedBy   string
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP()"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP()"`
	ServerId    string
}
