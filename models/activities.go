package models

import (
	"baby-varkin-health/initializers"
	"time"

	"gorm.io/gorm"
)

// ActityType

// drinkmilk
// pampers
// sleep
// wakeup
// takeabath
// pumpmilk
// milkmilk

type Activities struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	ActityType  string `json:"actity_type"`
	ActityValue string `json:"actity_value"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

// Save user details
func (activity *Activities) Save() (*Activities, error) {
	initializers.DB.Create(&activity)
	return activity, nil
}
