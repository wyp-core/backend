package repository

import "github.com/Abhyuday04/wyp/layers/models"

type IRepository interface {
	AddUser(fetchParams *models.AddUserFetchParam) (*models.AddUserRespParams, error)
	AddJob(job *models.Job) (*models.AddJobRespParams, error)
}