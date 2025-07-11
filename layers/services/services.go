package services

import (
	// "errors"

	"context"

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

func (s *Service) AddUser(fetchParams *models.AddUserFetchParam) (*models.AddUserRespParams, error) {
	// Call the repository method to add a user
	respParams, err := s.repository.AddUser(fetchParams)
	if err != nil {
		return nil, err
	}
	return respParams, nil
}

func (s *Service) AddJob(job *models.Job) (*models.AddJobRespParams, error) {
	// Call the repository method to add a user
	respParams, err := s.repository.AddJob(job)
	if err != nil {
		return nil, err
	}
	return respParams, nil
}

func (s *Service) GetJobs(fetchParams *models.GetJobsFetchParam) ([]models.Job, error) {

	// if fetchParams.UserID != "" {
	// 	jobs, err := s.repository.GetJobsByUserID(fetchParams)
	// 	if err != nil {
	// 		return nil, errors.New("failed to fetch jobs by user ID: " + err.Error())
	// 	}
	// 	return jobs, nil
	// }

	jobs, err := s.repository.GetJobs(fetchParams)
	if err != nil {
		return nil, err
	}
	return jobs, nil
}

func (s *Service) SendOtp(ctx context.Context, fetchParams *models.SendOtpParam) error {
	var otpStruct models.OtpRedisStruct
	otpStruct.CountryCode = fetchParams.CountryCode
	otpStruct.Phone = fetchParams.Phone
	err := otpStruct.GenerateOtp(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) VerifyOtp(ctx context.Context, fetchParams *models.VerifyOtpParam) (string, error) {
	var otpStruct models.OtpRedisStruct
	otpStruct.CountryCode = fetchParams.CountryCode
	otpStruct.Phone = fetchParams.Phone
	otpStruct.Otp = fetchParams.Otp

	err := otpStruct.IncrementAttemptCount(ctx)
	if err != nil {
		return "", err
	}

	userID, err := otpStruct.VerifyOtp(ctx)
	if err != nil {
		return "", err
	}
	return userID, nil
}