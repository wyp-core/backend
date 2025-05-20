package services

import "github.com/Abhyuday04/wyp/layers/models"


type IService interface {
	AddUser(fetchParams *models.AddUserFetchParam) (*models.AddUserRespParams, error)
}