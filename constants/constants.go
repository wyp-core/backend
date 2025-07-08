package constants

const (
	ErrorInvalidRequestBody = "Invalid request body"
)

const (
	OtpExpirationTimeMin = 5 // 5 minutes
	OtpMaxResend = 3
	ErrorOtpMaxResendExceeded = "maximum OTP attempts exceeded"
	MessageOtpMatch = "matched"
	MessageOtpNotMatched = "not matched"
)
