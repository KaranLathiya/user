package dal

import (
	"database/sql"
	error_handling "user/error"
	"user/model/request"
)

func CreateUser(db *sql.DB, verifyOTP request.VerifyOTP) (string, error) {
	var userID string
	err := db.QueryRow("select unique_rowid();").Scan(&userID)
	if err != nil {
		return "", error_handling.InternalServerError
	}
	err = db.QueryRow("INSERT INTO public.users (id, firstname, lastname, fullname, username, email, phone_number, country_code, signup_mode) VALUES ( $1, $2, $3, $4, $5, $6, $7, $8, $9 ) returning id;", userID, verifyOTP.Firstname, verifyOTP.Lastname, verifyOTP.Firstname+" "+verifyOTP.Lastname, userID, verifyOTP.Email, verifyOTP.PhoneNumber, verifyOTP.CountryCode, verifyOTP.SignupMode).Scan(&userID)
	if err != nil {
		return "", error_handling.DatabaseErrorShow(err)
	}
	return userID, nil
}
