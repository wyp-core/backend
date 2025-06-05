package models

import (
	"time"
)

type User struct {
	UserID      string    `json:"userID" gorm:"primary_key;column:user_id;type:uuid;default:uuid_generate_v4()"`
	Name        string    `json:"name" gorm:"size:100;not null"`
	Phone       string    `json:"phone" gorm:"size:20;not null"`
	CountryCode string    `json:"countryCode" gorm:"column:country_code;size:5;not null"`
	Age         int       `json:"age" gorm:"not null"`
	Gender      string    `json:"gender" gorm:"type:gender_type;not null"`
	CreatedAt   time.Time `json:"createdAt" gorm:"column:created_at;autoCreateTime"`
}

type Job struct {
	JobID          string    `json:"jobID" gorm:"primary_key;column:job_id;type:uuid;default:uuid_generate_v4()"`
	CreatedBy      string    `json:"createdBy" gorm:"column:created_by;type:uuid;not null"`
	Title          string    `json:"title" gorm:"type:text;not null"`
	Description    string    `json:"description" gorm:"type:text;not null"`
	Lat            float64   `json:"lat" gorm:"not null"`
	Lon            float64   `json:"lon" gorm:"not null"`
	CreatedAt      time.Time `json:"createdAt" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt      time.Time `json:"updatedAt" gorm:"column:updated_at;autoUpdateTime"`
	IsActive       bool      `json:"isActive" gorm:"column:is_active;default:true"`
	Price          float64   `json:"price" gorm:"not null"`
	Category       string    `json:"category" gorm:"type:text"`
	Mode           string    `json:"mode" gorm:"type:mode_type;not null"`
	Views          int       `json:"views" gorm:"default:0"`
	Duration       string    `json:"duration" gorm:"type:text"`
	GeoLocation    string    `json:"-" gorm:"type:geometry"`
	DistanceMeters float64   `json:"distanceMeters" gorm:"-"`
}
