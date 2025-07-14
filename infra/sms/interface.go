package sms

import "context"


type ISms interface {	
	SendOtp(ctx context.Context, countryCode string, phone string, otp int) error
}