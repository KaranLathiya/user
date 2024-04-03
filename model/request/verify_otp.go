package request

type VerifyOTP struct {
	Email       string `json:"email" validate:"email,required_without=phoneNumber" `
	EventType   string `json:"eventType" validate:"required,oneof=signup login organization_delete" `
	PhoneNumber string `json:"phoneNumber" validate:"len=10,required_without=email"   `
	CountryCode string `json:"countryCode" validate:"required_with=phoneNumber" `
	OTP         string `json:"otp" validate:"required,len=6" `
}
