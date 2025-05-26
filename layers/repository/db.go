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

func (r *Repository) GetJobs(fetchParams *models.GetJobsFetchParam) ([]models.Job, error) {
	var jobs []models.Job
	
	// Start building the query
	query := r.DatabaseClient.Model(&models.Job{})
	
	// Filter by UserID if provided
	if fetchParams.UserID != "" {
		query = query.Where("user_id = ?", fetchParams.UserID)
	}
	
	// Filter by price range
	if fetchParams.MinPrice > 0 {
		query = query.Where("price >= ?", fetchParams.MinPrice)
	}
	if fetchParams.MaxPrice > 0 {
		query = query.Where("price <= ?", fetchParams.MaxPrice)
	}
	
	// Filter by mode if provided
	if fetchParams.Mode != "" {
		query = query.Where("mode = ?", fetchParams.Mode)
	}
	
	// Handle sorting
	orderClause := "created_at DESC" // default sort
	if fetchParams.SortBy != "" {
		switch fetchParams.SortBy {
		case "price_asc":
			orderClause = "price ASC"
		case "price_desc":
			orderClause = "price DESC"
		case "createdAt_desc", "created_at":
			orderClause = "created_at DESC"
		case "createdAt_asc", "created_at_asc":
			orderClause = "created_at ASC"
		default:
			orderClause = "created_at DESC"
		}
	}
	query = query.Order(orderClause)
	
	// Handle pagination
	if fetchParams.Limit > 0 {
		query = query.Limit(fetchParams.Limit)
	}
	if fetchParams.Offset > 0 {
		query = query.Offset(fetchParams.Offset)
	}
	
	// Execute the query
	if err := query.Find(&jobs).Error; err != nil {
		return nil, err
	}
	
	return jobs, nil
}
