package models

type Users struct {
	UserID      string `json:"userID" db:"user_id"`
	Name        string `json:"name" db:"name"`
	Phone       string `json:"phone" db:"phone"`
	CountryCode string `json:"countryCode" db:"country_code"`
	Age         int    `json:"age" db:"age" `
	Gender      string `json:"gender" db:"gender" `
}
