package services

import "github.com/Abhyuday04/wyp/layers/models"


type IService interface {
	AddUser(fetchParams *models.AddUserFetchParam) (*models.AddUserRespParams, error)
	AddJob(job *models.Job) (*models.AddJobRespParams, error)
	GetJobs(job *models.GetJobsFetchParam) ([]models.Job, error)
}