package services

import (
	"github.com/Abhyuday04/wyp/layers/models"
	repositoryuser "github.com/Abhyuday04/wyp/layers/repository/users"
)

type Service struct {
	repositoryUser repositoryuser.IUserRepository
}
func New(repositoryUser repositoryuser.IUserRepository) *Service {
	return &Service{
		repositoryUser: repositoryUser,
	}
}

// AddUser implements IService.
func (s *Service) AddUser(fetchParams *models.AddUserFetchParam) (*models.AddUserRespParams, error) {
	// Call the repository method to add a user
	respParams ,err := s.repositoryUser.AddUser(fetchParams)
	if err != nil {
		return nil, err
	}
	return respParams, nil
}

