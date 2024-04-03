package dal

import (
	"database/sql"
	"fmt"
	error_handling "user/error"
	"user/model/request"
	"user/utils"
)

func StoreOTP(db *sql.DB, signup request.Signup, otp string, eventType string) error {
	var err error
	if signup.EventType == "email" {
		_, err = db.Exec("INSERT INTO public.otp (email,otp,expires_at,event_type) VALUES ($1,$2,$3,$4)", signup.Email, otp, utils.CurrentUTCTime(5), eventType)
	} else {
		_, err = db.Exec("INSERT INTO public.otp (phone_number,country_code,otp,expires_at,event_type) VALUES ($1,$2,$3,$4,$5)", signup.PhoneNumber, signup.CountryCode, otp, utils.CurrentUTCTime(5), eventType)
	}
	fmt.Println(err)
	if err != nil {
		return error_handling.InternalServerError
	}
	return nil
}
