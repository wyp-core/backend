package repositoryuser

import (
	"fmt"
	"github.com/Abhyuday04/wyp/layers/models"
	"gorm.io/gorm"
	"github.com/rs/zerolog/log"
)

type UserRepository struct {
	DatabaseClient *gorm.DB
}

func New(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DatabaseClient: db,
	}
}

// AddUser adds a new user to the database
func (r *UserRepository) AddUser(fetchParams *models.AddUserFetchParam) (*models.AddUserRespParams, error) {
	var response models.AddUserRespParams
	fmt.Print(fetchParams)
	user := models.User{
		Name:        fetchParams.Name,
		Phone:       fetchParams.Phone,
		CountryCode: fetchParams.CountryCode,
		Age:         fetchParams.Age,
		Gender:      fetchParams.Gender,
	}
	result := r.DatabaseClient.Create(&user)
	if result.Error != nil {
		log.Error().Err(result.Error).Msg("Failed to insert user")
		return &models.AddUserRespParams{}, fmt.Errorf("failed to insert user: %w", result.Error)
	}
	response.UserID = user.UserID
	response.Token = ""

	return &response, nil
}
