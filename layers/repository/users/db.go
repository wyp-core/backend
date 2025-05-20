package repositoryuser

import (
	"fmt"

	"github.com/Abhyuday04/wyp/layers/models"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	DatabaseClient *sqlx.DB
}

func New(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		DatabaseClient: db,
	}
}

// AddUser adds a new user to the database
func (r *UserRepository) AddUser(fetchParams *models.AddUserFetchParam) (*models.AddUserRespParams, error) {
	query := `
		INSERT INTO users (name, phone, country_code, age, gender)
		VALUES ($1, $2, $3, $4, $5)
	`
	var response models.AddUserRespParams
	fmt.Print(fetchParams)
	// Execute query and scan the returned user_id directly into the response struct
	// err := r.DatabaseClient.QueryRow(
	// 	query,
	// 	fetchParams.Name,
	// 	fetchParams.Phone,
	// 	fetchParams.CountryCode,
	// 	fetchParams.Age,
	// 	fetchParams.Gender,
	// ).Scan(&response.UserID)

	fmt.Println("CHECKING DB CLIENT IN REPO:", r)

	_, err := r.DatabaseClient.Exec(query, fetchParams.Name, fetchParams.Phone, fetchParams.CountryCode, fetchParams.Age, fetchParams.Gender)
	if err != nil {
		return &models.AddUserRespParams{}, fmt.Errorf("failed to insert user: %w", err)
	}
	// fmt.Println("User ID:", userID)

	// Set token field (if needed, replace with your token generation logic)
	response.Token = ""

	fmt.Println("User inserted successfully!")

	return &response, nil
}
