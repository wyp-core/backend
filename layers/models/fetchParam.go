package models

type AddUserFetchParam struct {
	Name        string `json:"name"`
	Age         int    `json:"age"`
	Gender      string `json:"gender"`
	Phone       string `json:"phone"`
	CountryCode string `json:"countryCode"`
}

