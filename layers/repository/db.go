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
    
    err := r.DatabaseClient.Transaction(func(tx *gorm.DB) error {
        // Create job without geo_location field
        if err := tx.Omit("geo_location").Create(job).Error; err != nil {
            return err
        }
        
        // Update geo_location separately
        if err := tx.Model(job).Update("geo_location", 
            gorm.Expr("ST_SetSRID(ST_MakePoint(?, ?), 4326)", job.Lon, job.Lat)).Error; err != nil {
            return err
        }
        
        return nil
    })
    
    if err != nil {
        log.Error().Err(err).Msg("Failed to insert job")
        return &models.AddJobRespParams{}, fmt.Errorf("failed to insert job: %w", err)
    }
    
    response.JobID = job.JobID
    return &response, nil
}

func (r *Repository) GetJobs(fetchParams *models.GetJobsFetchParam) ([]models.Job, error) {
	var jobs []models.Job
	
	// Start building the query
	query := r.DatabaseClient.Model(&models.Job{})
	query = query.Select("jobs.*, ST_Distance(geo_location, CAST(ST_SetSRID(ST_MakePoint(?, ?), 4326) AS geography)) AS distance_meters ",fetchParams.Lon, fetchParams.Lat)
	
	// Filter by UserID if provided
	if fetchParams.UserID != "" {
		query = query.Where("created_by = ?", fetchParams.UserID)
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

	if fetchParams.Radius != 0 {
		query = query.Where("ST_DWithin(geo_location, CAST(ST_SetSRID(ST_MakePoint(?, ?), 4326) AS geography), ?)", 
			fetchParams.Lon, fetchParams.Lat, fetchParams.Radius)
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
		case "radius_asc":
			orderClause = "distance_meters ASC"
		case "radius_desc":
			orderClause = "distance_meters DESC"
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
