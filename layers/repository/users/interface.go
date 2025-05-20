package repositoryuser

import "github.com/Abhyuday04/wyp/layers/models"

type IUserRepository interface {
	AddUser(fetchParams *models.AddUserFetchParam) (*models.AddUserRespParams, error)
}