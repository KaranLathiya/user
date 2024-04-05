package dal

import (
	"database/sql"
	"fmt"
	error_handling "user/error"
	"user/model/request"
	"user/utils"

	"github.com/lib/pq"
)

func UserCreate(db *sql.DB, verifyOTP request.VerifyOTP) (string, error) {
	// var filterArgsList []interface{}
	var query string
	var userID string
	err := db.QueryRow("select unique_rowid();").Scan(&userID)
	if err != nil {
		return "", error_handling.InternalServerError
	}
	if verifyOTP.SignupMode == "email" || verifyOTP.SignupMode == "google_login" {
		query = fmt.Sprintf("INSERT INTO public.users (id, firstname, lastname, fullname, username, email, created_at, signup_mode) VALUES ('%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v') returning id;", userID, verifyOTP.Firstname, verifyOTP.Lastname, verifyOTP.Firstname+" "+verifyOTP.Lastname, userID, verifyOTP.Email, utils.CurrentUTCTime(0), verifyOTP.SignupMode)
		// filterArgsList = append(filterArgsList, verifyOTP.Email)
	} else if verifyOTP.SignupMode == "phone" {
		query = fmt.Sprintf("INSERT INTO public.users (id, firstname, lastname, fullname, username, phone_number, country_code, created_at, signup_mode) VALUES ('%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v') returning id;", userID, verifyOTP.Firstname, verifyOTP.Lastname, verifyOTP.Firstname+" "+verifyOTP.Lastname, userID, verifyOTP.PhoneNumber, verifyOTP.CountryCode, utils.CurrentUTCTime(0), verifyOTP.SignupMode)
		// filterArgsList = append(filterArgsList, verifyOTP.PhoneNumber, verifyOTP.CountryCode)
	}
	err = db.QueryRow(query).Scan(&userID)
	if err != nil {
		if dbErr, ok := err.(*pq.Error); ok {
			errCode := dbErr.Code
			switch errCode {
			case "23505":
				// unique constraint violation
				return "", error_handling.UserAlreadyExist
			}
			return "", error_handling.InternalServerError
		}
	}
	return userID, nil
}
