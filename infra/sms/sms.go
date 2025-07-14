package sms

import (
	"context"
	"strconv"

	"github.com/rs/zerolog/log"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

type Sms struct {
	SmsClient *twilio.RestClient
}

func New(s *twilio.RestClient) *Sms {
	return &Sms{
		SmsClient: s,
	}
}

func (s *Sms) SendOtp(ctx context.Context, countryCode string, phone string, otp int) error {
	var number = ("+"+countryCode + phone)
	msg := "the otp is " + strconv.Itoa(otp)
	params := &openapi.CreateMessageParams{}
	params.SetBody(msg)
	params.SetFrom("+13134748210")
	params.SetTo(number)
	_, err := s.SmsClient.Api.CreateMessage(params)
	if err != nil {
		log.Error().Err(err).Msg("failed to send OTP")
		return err
	}
	return nil // Placeholder for actual implementation
}
