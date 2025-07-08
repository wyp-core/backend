package services

import (
	"context"

	"github.com/Abhyuday04/wyp/layers/models"
)


type IService interface {
	AddUser(fetchParams *models.AddUserFetchParam) (*models.AddUserRespParams, error)
	AddJob(job *models.Job) (*models.AddJobRespParams, error)
	GetJobs(job *models.GetJobsFetchParam) ([]models.Job, error)
	SendOtp(ctx context.Context, fetchParams *models.SendOtpParam) (error)
	VerifyOtp(ctx context.Context, fetchParams *models.VerifyOtpParam) (string, error)
}