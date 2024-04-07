package dal

import (
	"database/sql"
	"fmt"
	error_handling "user/error"
	"user/model/request"
	"user/utils"
)

func StoreOTP(db *sql.DB, storeOTP request.StoreOTP) error {
	var err error
	if storeOTP.LoginType == "email" || storeOTP.LoginType == "google_login"{
		_, err = db.Exec("INSERT INTO public.otp (email,otp,expires_at,event_type) VALUES ($1,$2,$3,$4)", storeOTP.Email, storeOTP.HashedOTP, utils.CurrentUTCTime(5), storeOTP.EventType)
	} else {
		_, err = db.Exec("INSERT INTO public.otp (phone_number,country_code,otp,expires_at,event_type) VALUES ($1,$2,$3,$4,$5)", storeOTP.PhoneNumber, storeOTP.CountryCode, storeOTP.HashedOTP, utils.CurrentUTCTime(5), storeOTP.EventType)
	}
	fmt.Println(err)
	if err != nil {
		return error_handling.InternalServerError
	}
	return nil
}
