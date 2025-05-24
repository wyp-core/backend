package services

import (
	"github.com/Abhyuday04/wyp/layers/models"
	repository "github.com/Abhyuday04/wyp/layers/repository"
)

type Service struct {
	repository repository.IRepository
}
func New(repository repository.IRepository) *Service {
	return &Service{
		repository: repository,
	}
}

// AddUser implements IService.
func (s *Service) AddUser(fetchParams *models.AddUserFetchParam) (*models.AddUserRespParams, error) {
	// Call the repository method to add a user
	respParams ,err := s.repository.AddUser(fetchParams)
	if err != nil {
		return nil, err
	}
	return respParams, nil
}

// AddJob implements IService.
func (s *Service) AddJob(job *models.Job) (*models.AddJobRespParams, error) {
	// Call the repository method to add a user
	respParams ,err := s.repository.AddJob(job)
	if err != nil {
		return nil, err
	}
	return respParams, nil
}

