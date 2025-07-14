package models

type AddUserFetchParam struct {
	Name        string `json:"name"`
	Age         int    `json:"age"`
	Gender      string `json:"gender"`
	Phone       string `json:"phone"`
	CountryCode string `json:"countryCode"`
}

type GetJobsFetchParam struct {
	UserID   string  `json:"userId"`
	MinPrice float64 `json:"minPrice"`
	MaxPrice float64 `json:"maxPrice"`
	Mode     string  `json:"mode"`
	Radius   float64 `json:"radius"`
	SortBy   string  `json:"sortBy"`
	Limit    int     `json:"limit"`
	Offset   int     `json:"offset"`
	Lat      float64 `json:"lat"`
	Lon      float64 `json:"lon"`
}

type SendOtpParam struct {
	Phone       string `json:"phone" required:"true"`
	CountryCode string `json:"countryCode" required:"true"`
}

type VerifyOtpParam struct {
	Phone       string `json:"phone"`
	CountryCode string `json:"countryCode"`
	Otp         int `json:"otp"`
}
