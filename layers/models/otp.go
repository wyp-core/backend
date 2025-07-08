package models

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Abhyuday04/wyp/constants"
	"github.com/Abhyuday04/wyp/general"
	"github.com/Abhyuday04/wyp/infra/redis"
)

type OtpRedisStruct struct {
	Phone       string `json:"phone"`
	CountryCode string `json:"countryCode"`
	Otp 	   int    `json:"otp"`
	AttemptCount int    `json:"attemptCount"` // Number of attempts made to verify the OTP
}

func (s *OtpRedisStruct) GenerateOtp(ctx context.Context) (error){
	otp, err := general.GenerateRandom4DigitSecure()
	if err != nil {
		return fmt.Errorf("failed to generate OTP: %w", err)
	}
	s.Otp = otp
	redisKey:=fmt.Sprintf("redis_%s%s", s.CountryCode, s.Phone)
	jsonData, err := json.Marshal(s)
	if err != nil {
		return fmt.Errorf("failed to marshal OTP data: %w", err)
	}
	if err := redis.RedisClient.Set(ctx, redisKey, jsonData, constants.OtpExpirationTimeMin * time.Minute).Err(); err != nil {
		fmt.Println("Error setting OTP in Redis:", err)
		return fmt.Errorf("failed to set OTP in Redis: %w", err)
	}
	return nil
}

func (s *OtpRedisStruct) IncrementAttemptCount(ctx context.Context) error {
	redisKey := fmt.Sprintf("redis_%s%s", s.CountryCode, s.Phone)
	val, err := redis.RedisClient.Get(ctx, redisKey).Result()
	if err != nil {
		return fmt.Errorf("failed to get OTP from Redis: %w", err)
	}

	var otpData OtpRedisStruct
	if err := json.Unmarshal([]byte(val), &otpData); err != nil {
		return fmt.Errorf("failed to unmarshal OTP data: %w", err)
	}

	otpData.AttemptCount++
	if otpData.AttemptCount >= constants.OtpMaxResend {
		return fmt.Errorf(constants.ErrorOtpMaxResendExceeded)
	}

	jsonData, err := json.Marshal(otpData)
	if err != nil {
		return fmt.Errorf("failed to marshal OTP data: %w", err)
	}

	if err := redis.RedisClient.Set(ctx, redisKey, jsonData, constants.OtpExpirationTimeMin * time.Minute).Err(); err != nil {
		return fmt.Errorf("failed to update OTP attempt count in Redis: %w", err)
	}
	return nil
}

func (s *OtpRedisStruct) VerifyOtp (ctx context.Context) (string, error) {
	redisKey := fmt.Sprintf("redis_%s%s", s.CountryCode, s.Phone)
	val, err := redis.RedisClient.Get(context.Background(), redisKey).Result()
	if err != nil {
		return "", fmt.Errorf("failed to get OTP from Redis: %w", err)
	}

	var otpData OtpRedisStruct
	if err := json.Unmarshal([]byte(val), &otpData); err != nil {
		return "", fmt.Errorf("failed to unmarshal OTP data: %w", err)
	}

	if otpData.Otp == s.Otp {
		return constants.MessageOtpMatch, nil
	} else {
		return constants.MessageOtpNotMatched, nil
	}
}
