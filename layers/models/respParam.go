package models

type AddUserRespParams struct {
	UserID string `json:"userId" db:"user_id"`
	Token  string `json:"token"`
}

type AddJobRespParams struct {
	JobID string `json:"jobId" db:"job_id"`
}