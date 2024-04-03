package request

type RequestOTP struct {
	Email     string `json:"email" validate:"email,required_without=phoneNumber" `
	EventType string `json:"eventType" validate:"required,oneof=signup login organization_delete" `
	PhoneNumber string `json:"phoneNumber" validate:"len=,required_without=email"   `
	CountryCode string `json:"countryCode" validate:"required_with=phoneNumber" `
}