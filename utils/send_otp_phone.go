package utils

import (
	"user/config"

	error_handling "user/error"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

func SendOTPPhone(countryCode string, phoneNumber string, otp string, subject string) error {

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: config.ConfigVal.Twilio.AcountSID,
		Password: config.ConfigVal.Twilio.AuthToken,
	})

	params := &twilioApi.CreateMessageParams{}
	params.SetTo("+" + countryCode + phoneNumber)
	params.SetFrom(config.ConfigVal.Twilio.MessageFrom)
	params.SetBody(subject + otp)

	_, err := client.Api.CreateMessage(params)
	if err != nil {
		return error_handling.SendMessageError
	}
	return nil
}
