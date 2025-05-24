package repository

import (
	"fmt"
	"github.com/Abhyuday04/wyp/layers/models"
	"gorm.io/gorm"
	"github.com/rs/zerolog/log"
)

type Repository struct {
	DatabaseClient *gorm.DB
}

func New(db *gorm.DB) *Repository {
	return &Repository{
		DatabaseClient: db,
	}
}

// AddUser adds a new user to the database
// TODO make single struct passing
func (r *Repository) AddUser(fetchParams *models.AddUserFetchParam) (*models.AddUserRespParams, error) {
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

func (r *Repository) AddJob(job *models.Job) (*models.AddJobRespParams, error) {
	var response models.AddJobRespParams
	result := r.DatabaseClient.Create(job)
	if result.Error != nil {
		log.Error().Err(result.Error).Msg("Failed to insert job")
		return &models.AddJobRespParams{}, fmt.Errorf("failed to insert job: %w", result.Error)
	}
	response.JobID = job.JobID
	return &response, nil
}
