package models

import "time"

type User struct {
	UserID      string `json:"userID" gorm:"primary_key;column:user_id;type:uuid;default:uuid_generate_v4()"`
	Name        string `json:"name" gorm:"size:100;not null"`
	Phone       string `json:"phone" gorm:"size:20;not null"`
	CountryCode string `json:"countryCode" gorm:"column:country_code;size:5;not null"`
	Age         int    `json:"age" gorm:"not null"`
	Gender      string `json:"gender" gorm:"type:gender_type;not null"`
	CreatedAt   time.Time `json:"createdAt" gorm:"column:created_at;autoCreateTime"`
}
